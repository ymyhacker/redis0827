package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)


func (dm *DatabaseManager) Expire(key string, ttl int) CommandResponse {
	if ttl < 0 {
		return CommandResponse{Message: "Invalid TTL"}
	}

	_, exists := dm.Database.Get(key)
	if !exists {
		return CommandResponse{Message: "Key does not exist"}
	}

	expirationTime := time.Now().Add(time.Duration(ttl) * time.Second)
	dm.Database.SetExpiration(key, expirationTime)
	return CommandResponse{Message: "OK"}
}

func (db *Database) SetExpiration(key string, expiration time.Time) {
	db.expirations[key] = expiration
}
