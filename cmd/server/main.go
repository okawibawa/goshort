package main

import (
	"log"
	"net/http"

	"github.com/okawibawa/goshort/internal/database"
	"github.com/okawibawa/goshort/internal/handlers"
)

func main() {
	dbPool, err := database.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
		return
	}
	defer database.CloseDB(dbPool)

	handler := handlers.NewHandler(dbPool)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/shorten-url", handler.Shorten)

	log.Printf("listening and serving port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
