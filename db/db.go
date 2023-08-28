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
