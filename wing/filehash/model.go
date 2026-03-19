package filehash

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// FileHashMetadata represents the metadata of a file.
type FileHashMetadata struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	MD5  []byte `json:"md5"`
	SHA1 []byte `json:"sha1"`
	GCID []byte `json:"gcid"`
}

// Serialize serializes the metadata to a byte slice.
func (m *FileHashMetadata) Serialize() []byte {
	return []byte(fmt.Sprintf("name=%s|size=%d|md5=%x|sha1=%x|gcid=%x",
		m.Name, m.Size, m.MD5, m.SHA1, m.GCID))
}

// DeserializeMetadata deserializes the metadata from a byte slice.
func DeserializeMetadata(data []byte) (*FileHashMetadata, error) {
	s := string(data)
	m := &FileHashMetadata{}
	parts := strings.Split(s, "|")
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key, value := kv[0], kv[1]
		switch key {
		case "name":
			m.Name = value
		case "size":
			fmt.Sscanf(value, "%d", &m.Size)
		case "md5":
			m.MD5, _ = hex.DecodeString(value)
		case "sha1":
			m.SHA1, _ = hex.DecodeString(value)
		case "gcid":
			m.GCID, _ = hex.DecodeString(value)
		}
	}
	return m, nil
}

// ToHumanReadable returns the metadata in a human-readable format.
func (m *FileHashMetadata) ToHumanReadable() string {
	return fmt.Sprintf("Name: %s\nSize: %d\nMD5: %x\nSHA1: %x\nGCID: %x",
		m.Name, m.Size, m.MD5, m.SHA1, m.GCID)
}
