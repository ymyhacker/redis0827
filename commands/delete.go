package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)

func (db *Database) Delete(key string) bool {
	_, exists := db.data[key]
	if exists {
		delete(db.data, key)
		delete(db.expirations, key)
		return true
	}
	return false
}
