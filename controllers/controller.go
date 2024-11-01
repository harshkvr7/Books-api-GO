package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func GetAllBooks(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book

		if err := rows.Scan(&book.ID, &book.Name, &book.Rating, &book.Author, &book.Genre); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		books = append(books, book)

		fmt.Printf("Book: %+v\n", book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	rows, err := DB.Query("SELECT * FROM books WHERE id = $1", bookId)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	var books []Book

	for rows.Next() {
		var book Book

		if err := rows.Scan(&book.ID, &book.Name, &book.Rating, &book.Author, &book.Genre); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		books = append(books, book)

		fmt.Printf("Book: %+v\n", book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBook(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println("Error Decoding Json")
		return
	}

	fmt.Printf("Received Book: %+v\n", book)

	query := "INSERT INTO books (name, rating, author, genre) VALUES ($1, $2, $3, $4) RETURNING id"

	err := DB.QueryRow(query, book.Name, book.Rating, book.Author, book.Genre).Scan(&book.ID)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book added successfully",
		"book_id": book.ID,
	})
}

func UpdateBook(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	query := "UPDATE books SET name = $1, rating = $2, author = $3, genre = $4 WHERE id = $5"
	res, err := DB.Exec(query, book.Name, book.Rating, book.Author, book.Genre, bookId)
	if err != nil {
		fmt.Println("Error updating book:", err)
		return
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		fmt.Println("Book not found")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book updated successfully",
	})
}

func DeleteBook(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	query := "DELETE FROM books WHERE id = $1"
	res, err := DB.Exec(query, bookId)
	if err != nil {
		fmt.Println("Error deleting book:", err)
		return
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		fmt.Println("Book not found")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book deleted successfully",
	})
}
