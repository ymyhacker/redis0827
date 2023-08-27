package commands

import (
	"strconv"
	"time"

	"github.com/ymyhacker/redis0827/db"
)

type CommandResponse struct {
	Message string
}

type DatabaseManager struct {
	Database *db.Database
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
			return CommandResponse{Message: "Usage: SET key value"}
		}
		dm.Database.Set(args[0], args[1])
		return CommandResponse{Message: "OK"}
		}// ... (other cases)
	case "EXPIRE":{
		if len(args) != 2 {
			fmt.Println("Usage: EXPIRE key ttl")
			continue
		}

		ttl, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid TTL value")
				continue
			}

			response := dm.Expire(args[0], ttl)
			fmt.Println(response.Message)
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
	
	case "LIST":
		// keys := dm.ListKeys()
		keys := dm.Database.ListKeys()
		return CommandResponse{Message: fmt.Sprintf("Keys in the database: %v", keys)}
	default:
		return CommandResponse{Message: "Unknown command"}
	}
}

func (db *Database) Set(key, value string) {
	db.data[key] = value
}

// Add more functions for other commands...

func (db *Database) SetWithTTL(key, value string, ttl int) {
	db.data[key] = value
	if ttl > 0 {
		expirationTime := time.Now().Add(time.Duration(ttl) * time.Second)
		db.expirations[key] = expirationTime
	}
}

func (db *Database) ListKeys() []string {
	keys := []string{}
	for key := range db.data {
		keys = append(keys, key)
	}
	return keys
}


// ... (other functions)

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
