package main

import (
	"fmt"
	"log"

	"github.com/okawibawa/goshort/internal/database"
)

func main() {
	dbPool, err := database.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
		return
	}
	defer database.CloseDB(dbPool)

	fmt.Println("running")
}

// CREATE TABLE urls (
// 	id SERIAL PRIMARY KEY,
// 	original_url TEXT NOT NULL,
// 	shorten_url VARCHAR(8) UNIQUE NOT NULL,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );
