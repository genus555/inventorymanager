package database

import (
	"sync"
	"fmt"

	_ "modernc.org/sqlite"
)

var mu sync.Mutex

func (db *DB) checkCache(entry string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()

	entry_table := fmt.Sprintf("%s.%s", db.TableName, entry)
	if amount, ok := db.cache[entry_table]; ok {
		for i, cached := range db.cacheOrder {
			if cached == entry_table {
				db.cacheOrder = append(db.cacheOrder[:i], db.cacheOrder[i+1:]...)
				db.cacheOrder = append(db.cacheOrder, entry_table)
				break
			}
		}
		return amount, true
	}

	return 0, false
}

func (db *DB) addToCache(entry string, amount int) {
	mu.Lock()
	defer mu.Unlock()

	entry_table := fmt.Sprintf("%s.%s", db.TableName, entry)
	if _, cached := db.cache[entry_table]; !cached {
		db.cacheOrder = append(db.cacheOrder, entry_table)
	}
	db.cache[entry_table] = amount

	if len(db.cacheOrder) > db.cacheLimit {
		db.deleteFromCache()
	}
}

func (db *DB) deleteFromCache() {
	entry_table := db.cacheOrder[0]
	db.cacheOrder = db.cacheOrder[1:]
	delete(db.cache, entry_table)
}

func (db *DB) deleteEntryFromCache(entry string) {
	mu.Lock()
	defer mu.Unlock()
	entry_table := fmt.Sprintf("%s.%s", db.TableName, entry)
	for i, cached := range db.cacheOrder {
		if cached == entry_table {
			db.cacheOrder = append(db.cacheOrder[:i], db.cacheOrder[i+1:]...)
			break
		}
	}
	delete(db.cache, entry_table)
}