package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Database struct {
	*sql.DB
}

var db *Database
var once sync.Once

func DBconnection(host, user, password, port, dbname string) *Database {
	once.Do(func() {
		connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, dbname)

		dbConn, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		db = &Database{dbConn}
	})
	return db
}
