package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)
func (db *Database) KeyExists(key string) bool {
	_, exists := db.data[key]
	return exists
}
