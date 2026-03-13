package filehash

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMetadata(t *testing.T) {
	dir, err := os.MkdirTemp("", "hash-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.txt")
	content := []byte("hello openlist")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	m, err := CalculateMetadata(path)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, "test.txt", m.Name)
	assert.Equal(t, int64(len(content)), m.Size)
	assert.NotEmpty(t, m.MD5)
	assert.NotEmpty(t, m.SHA1)
	assert.NotEmpty(t, m.GCID)
}
