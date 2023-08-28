package commands

import (
	"strconv"
	"time"
	"fmt"
	"strings"
	// "github.com/ymyhacker/redis0827/tree/YmY-branch/db"
)


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


func (db *Database) Expire(key string, ttl int) CommandResponse {
	if ttl < 0 {
		return CommandResponse{Message: "Invalid TTL"}
	}

	_, exists := db.Get(key)
	if !exists {
		return CommandResponse{Message: "Key does not exist"}
	}

	expirationTime := time.Now().Add(time.Duration(ttl) * time.Second)
	db.SetExpiration(key, expirationTime)
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




func ExecuteCommand(dm *Database, command string, args []string) CommandResponse {
	
	fmt.Println(command)
	fmt.Println(i for i in args)
	switch strings.ToLower(command) {
		
	case "set":{
		if len(args) != 2 {
			return CommandResponse{Message: "Usage: SET key value expiretime"}
		}
		dm.Set(args[0], args[1])
		return CommandResponse{Message: "OK"}
		}// ... (other cases)
	case "expire":{
		if len(args) != 2 {
			return CommandResponse{Message: "Usage: EXPIRE key ttl"}
		}

		ttl, err := strconv.Atoi(args[1])
			if err != nil {
				return CommandResponse{Message: "Invalid TTL value"}
			}

			response := dm.Expire(args[0], ttl)
			return response

	}	

	case "delete":
		{
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: DELETE key"}
		}
		dm.Delete(args[0])
		return CommandResponse{Message: "OK"}
	}
	case "get":
		{
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: GET key"}
		}
		value,_ := dm.Get(args[0])
		return CommandResponse{Message: value}
		// ... (other cases)
		}
	case "exist":
		if len(args) != 1 {
			return CommandResponse{Message: "Usage: EXISTS key"}
		}
		key := args[0]
		if dm.KeyExists(key) {
			return CommandResponse{Message: fmt.Sprintf("Key %s exists", key)}
		}
		return CommandResponse{Message: fmt.Sprintf("Key %s does not exist", key)}
	case "list":
		// keys := dm.ListKeys()
		keys := dm.ListKeys()
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
