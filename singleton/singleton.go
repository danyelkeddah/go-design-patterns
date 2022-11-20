package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Database struct{}

var databaseInstance *Database

func GetDatabaseInstance() *Database {
	if databaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if databaseInstance == nil {
			fmt.Println("Creating database instance now")
			databaseInstance = &Database{}
		} else {
			fmt.Println("Single instance already created -1")
		}
	} else {
		fmt.Println("Single instance already created -2")
	}

	return databaseInstance
}
