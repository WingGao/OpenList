package filehash

import (
	"crypto/md5"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"

	hashutil "github.com/OpenListTeam/OpenList/v4/pkg/utils/hash"
)

// CalculateHashMetadata calculates the metadata for a file.
func CalculateHashMetadata(path string) (*FileHashMetadata, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, nil
	}

	hMD5 := md5.New()
	hSHA1 := sha1.New()
	hGCID := hashutil.NewGcid(info.Size())

	mw := io.MultiWriter(hMD5, hSHA1, hGCID)
	if _, err := io.Copy(mw, f); err != nil {
		return nil, err
	}

	return &FileHashMetadata{
		Name: filepath.Base(path),
		Size: info.Size(),
		MD5:  hMD5.Sum(nil),
		SHA1: hSHA1.Sum(nil),
		GCID: hGCID.Sum(nil),
	}, nil
}
