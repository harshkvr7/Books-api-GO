package main

import (
	"books/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Books API")

	connStr := os.Getenv("DATABASE_URL")

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer DB.Close()

	BooksRouter := router.Router(DB)

	http.ListenAndServe(":3000", BooksRouter)
}
