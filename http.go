package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID     string  `json:"id"`
	Author *Author `json:"author"`
	Title  string  `json:"title"`
	
}

type Author struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

var books []Book

func getBooks(w http.ResponseWriter, w *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, w *http.Request) {

}

func createBook(w http.ResponseWriter, w *http.Request) {

}

func updateBook(w http.ResponseWriter, w *http.Request) {

}

func deleteBook(w http.ResponseWriter, w *http.Request) {

}

func main() {
	reg := mux.NewRouter()

	books = append(books, Book{ID: "1", Author: &Author{Firstname: "John", Lastname: "Doe"}, Title: "Book One"})
	books = append(books, Book{ID: "2", Author: &Author{Firstname: "Thomas", Lastname: "Steve"}, Title: "Book Two"})
	books = append(books, Book{ID: "3", Author: &Author{Firstname: "Jack", Lastname: "Andrew"}, Title: "Book Three"})
	books = append(books, Book{ID: "4", Author: &Author{Firstname: "Tim", Lastname: "C"}, Title: "Book Four"})

	r.HandleFunc("/api/books", getBooks).methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).methods("GET")
	r.HandleFunc("/api/books", createBook).methods("PUSH")
	r.HandleFunc("/api/books/{id}", updateBook).methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}