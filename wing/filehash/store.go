package filehash

import (
	"fmt"

	"github.com/cockroachdb/pebble"
)

const (
	PrefixMD5  = "K:MD5:"
	PrefixSHA1 = "K:SHA1:"
	PrefixGCID = "K:GCID:"
	PrefixPath = "C:"
)

// SaveMetadata saves the metadata to the database.
func SaveMetadata(m *FileHashMetadata, path string) error {
	batch := db.NewBatch()
	defer batch.Close()

	// 1. Primary MD5 key
	md5Key := []byte(PrefixMD5 + string(m.MD5))
	if err := batch.Set(md5Key, m.Serialize(), pebble.Sync); err != nil {
		return err
	}

	// 2. SHA1 lookup key
	sha1Key := []byte(PrefixSHA1 + string(m.SHA1))
	if err := batch.Set(sha1Key, m.MD5, pebble.Sync); err != nil {
		return err
	}

	// 3. GCID lookup key
	if len(m.GCID) > 0 {
		gcidKey := []byte(PrefixGCID + string(m.GCID))
		if err := batch.Set(gcidKey, m.MD5, pebble.Sync); err != nil {
			return err
		}
	}

	// 4. Local path/size cache key
	if path != "" {
		pathKey := []byte(fmt.Sprintf("%s%s|%d", PrefixPath, path, m.Size))
		if err := batch.Set(pathKey, m.MD5, pebble.Sync); err != nil {
			return err
		}
	}

	return batch.Commit(pebble.Sync)
}

// GetMetadataByMD5 retrieves metadata by MD5.
func GetMetadataByMD5(md5 []byte) (*FileHashMetadata, error) {
	key := []byte(PrefixMD5 + string(md5))
	val, closer, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	return DeserializeMetadata(val)
}

// GetMetadataBySHA1 retrieves metadata by SHA1.
func GetMetadataBySHA1(sha1 []byte) (*FileHashMetadata, error) {
	key := []byte(PrefixSHA1 + string(sha1))
	md5, closer, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	return GetMetadataByMD5(md5)
}

// GetMetadataByGCID retrieves metadata by GCID.
func GetMetadataByGCID(gcid []byte) (*FileHashMetadata, error) {
	key := []byte(PrefixGCID + string(gcid))
	md5, closer, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	return GetMetadataByMD5(md5)
}

// GetMetadataByPath retrieves metadata by path and size.
func GetMetadataByPath(path string, size int64) (*FileHashMetadata, error) {
	key := []byte(fmt.Sprintf("%s%s|%d", PrefixPath, path, size))
	md5, closer, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	return GetMetadataByMD5(md5)
}
