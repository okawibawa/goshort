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
