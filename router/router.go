package router

import (
	"books/controllers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(DB *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllBooks(DB, w, r)
	}).Methods("GET")

	router.HandleFunc("/api/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetBook(DB, w, r)
	}).Methods("GET")

	router.HandleFunc("/api/book", func(w http.ResponseWriter, r *http.Request) {
		controllers.AddBook(DB, w, r)
	}).Methods("POST")

	router.HandleFunc("/api/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateBook(DB, w, r)
	}).Methods("PUT")

	router.HandleFunc("/api/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteBook(DB, w, r)
	}).Methods("DELETE")

	return router
}
