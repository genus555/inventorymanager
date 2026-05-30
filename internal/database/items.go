package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	TableName	string
	cache		map[string]int
	cacheOrder	[]string
	cacheLimit	int
}

const (
	PLUS = true
	MINUS = false
)

const (
	LOW = 1
	DEFAULTAMOUNT = 0
	STARTCACHELIMIT = 5
)

func NewDB(db *sql.DB) *DB {
	return &DB{
		database:	db,
		TableName:	"",
		cache:		make(map[string]int),
		cacheOrder:	[]string{},
		cacheLimit:	STARTCACHELIMIT,
	}
}