package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating singleton for database")
	time.Sleep(3 * time.Second)
	fmt.Println("Singleton for database created")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating a new instance of a database connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("**DB ALREDY CREATED***")
	}
	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()

}
