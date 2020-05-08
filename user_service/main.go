package main

import (
	"log"
)

func main() {

	db, err := createPostgresDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Println("Connected to DB successfully")
	}
}
