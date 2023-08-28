package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)

func (db *Database) ListKeys() []string {
	keys := []string{}
	for key := range db.data {
		keys = append(keys, key)
	}
	return keys
}
