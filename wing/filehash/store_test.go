package filehash

import (
	"os"
	"testing"

	"github.com/cockroachdb/pebble"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) {
	dir, err := os.MkdirTemp("", "filehash-test")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })

	db, err = pebble.Open(dir, &pebble.Options{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
}

func TestStore(t *testing.T) {
	setupTestDB(t)

	m := &FileHashMetadata{
		Name: "test.txt",
		Size: 100,
		MD5:  []byte("md5hash"),
		SHA1: []byte("sha1hash"),
		GCID: []byte("gcidhash"),
	}

	path := "/abs/path/test.txt"

	// Test Save
	err := SaveMetadata(m, path)
	assert.NoError(t, err)

	// Test Get by MD5
	m2, err := GetMetadataByMD5(m.MD5)
	assert.NoError(t, err)
	assert.Equal(t, m.Name, m2.Name)
	assert.Equal(t, m.Size, m2.Size)

	// Test Get by SHA1
	m3, err := GetMetadataBySHA1(m.SHA1)
	assert.NoError(t, err)
	assert.Equal(t, m.MD5, m3.MD5)

	// Test Get by GCID
	m4, err := GetMetadataByGCID(m.GCID)
	assert.NoError(t, err)
	assert.Equal(t, m.MD5, m4.MD5)

	// Test Get by Path
	m5, err := GetMetadataByPath(path, 100)
	assert.NoError(t, err)
	assert.Equal(t, m.MD5, m5.MD5)

	// Test Get by Path (wrong size)
	_, err = GetMetadataByPath(path, 101)
	assert.Error(t, err)
}
