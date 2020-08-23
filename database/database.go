package database

import (
	"fmt"
)

type Database interface {
	// 2. Patch database: update current entries if exists, otherwise insert
	// 2-1. If exists: Compare update timestamp and set is_viewed to false if updated
	// 3. Get all jobs from database
	List(data interface{}) error
	Patch(key interface{}, value []byte)
	Insert(key interface{}, value []byte)
	Commit() error
}

func Open(dbname string) (Database, error) {
	filename := fmt.Sprintf("./%s.json", dbname)
	return newJSONDatabase(filename)
}
