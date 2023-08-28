package commands

import (
	"strconv"
	"time"
	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)

func (db *Database) Set(key, value string)  {
	db.data[key] = value
    da.expiresAt = time.Now().Add(60 * time.Second)
}
