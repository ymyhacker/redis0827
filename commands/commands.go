package commands

import (
	"strconv"
	"time"
	"github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)
package db

import (
	   "time"
	//    "strings"
)
type DatabaseManager struct {
	Database *Database
}


type CommandResponse struct {
	Message string
}
type Database struct {
	data        map[string]string
	expirations map[string]time.Time
}

func NewDatabase() *Database {
	return &Database{
		data:        make(map[string]string),
		expirations: make(map[string]time.Time),
	}
}

func (db *Database) Set(key, value string)  {
	db.data[key] = value
    db.expirations[key] = time.Now().Add(60 * time.Second)
}

func (db *Database) Delete(key string) bool {
	_, exists := db.data[key]
	if exists {
		delete(db.data, key)
		delete(db.expirations, key)
		return true
	}
	return false


}

func (db *Database) KeyExists(key string) bool {
	_, exists := db.data[key]
	return exists
}


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

func (db *Database) ListKeys() []string {
	keys := []string{}
	for key := range db.data {
		keys = append(keys, key)
	}
	return keys
}

type CommandResponse struct {
	Message string
}

type DatabaseManager struct {
	Database *Database
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		Database: db.NewDatabase(),
	}
}

func InitDatabase() *DatabaseManager {
	return &DatabaseManager{
		Database: db.NewDatabase(),
	}
}

func ExecuteCommand(dm *DatabaseManager, command string, args []string) CommandResponse {
	switch command {
	case "SET":{
		if len(args) != 2 {
			return CommandResponse{Message: "Usage: SET key value expiretime"}
		}
		dm.Database.Set(args[0], args[1])
		return CommandResponse{Message: "OK"}
		}// ... (other cases)
	case "EXPIRE":{
		if len(args) != 2 {
			return CommandResponse{Message: "Usage: EXPIRE key ttl"}
		}

		ttl, err := strconv.Atoi(args[1])
			if err != nil {
				return CommandResponse{Message: "Invalid TTL value"}
			}

			response := dm.Expire(args[0], ttl)
			return CommandResponse{Message: "%s",response.Message}

	}	

	case "DELETE":
		{
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: DELETE key"}
		}
		dm.Database.Delete(args[0])
		return CommandResponse{Message: "OK"}
	}
	case "GET":
		{
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: GET key"}
		}
		value := dm.Database.Get(args[0])
		return CommandResponse{Message: value}
		// ... (other cases)
		}
	case "Exist":
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: EXISTS key"}
		}
		key := args[0]
		if dm.Database.KeyExists(key) {
			return CommandResponse{Message: fmt.Sprintf("Key %s exists", key)}
		}
		return CommandResponse{Message: fmt.Sprintf("Key %s does not exist", key)}
	case "LIST":
		// keys := dm.ListKeys()
		keys := dm.Database.ListKeys()
		return CommandResponse{Message: fmt.Sprintf("Keys in the database: %v", keys)}
	default:
		return CommandResponse{Message: "Unknown command"}
	}
}



// Add more functions for other commands...

func (db *Database) SetWithTTL(key, value string, ttl int) {
	db.data[key] = value
	if ttl > 0 {
		expirationTime := time.Now().Add(time.Duration(ttl) * time.Second)
		db.expirations[key] = expirationTime
	}
}



// ... (other functions)
