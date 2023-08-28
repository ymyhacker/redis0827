package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)

func (db *Database) Get(key string) (string, bool) {
	value, exists := db.data[key]
	if exists {
		expirationTime, hasExpiration := db.expirations[key]
		if hasExpiration && time.Now().After(expirationTime) {
			delete(db.data, key)
			delete(db.expirations, key)
			return "", false
		}
		return value, true
	}
	return "", false
}
