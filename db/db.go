package db

import (
	   "time"
	//    "strings"
)
// type Database struct {
//     data     map[string]string
//     expiresAt map[string]time.Duration
// }


// func NewDatabase() *Database {
// 	return &Database{
// 		data: make(map[string]string),
// 		expiresAt: make(map[string]time.Duration),
// 	}
// }


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
    da.expiresAt = time.Now().Add(60 * time.Second)
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

func (db *Database) Set(key, value string)  {
	db.data[key] = value
    da.expiresAt = time.Now().Add(60 * time.Second)
}
