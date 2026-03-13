package filehash

import (
	"path/filepath"
	"sync"

	"github.com/OpenListTeam/OpenList/v4/cmd/flags"
	"github.com/cockroachdb/pebble"
)

var (
	db   *pebble.DB
	once sync.Once
)

// Init initializes the metadata database.
func Init() error {
	var err error
	once.Do(func() {
		dir := filepath.Join(flags.DataDir, "filehash")
		db, err = pebble.Open(dir, &pebble.Options{})
	})
	return err
}

// Close closes the metadata database.
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// GetDB returns the underlying pebble database instance.
func GetDB() *pebble.DB {
	return db
}
