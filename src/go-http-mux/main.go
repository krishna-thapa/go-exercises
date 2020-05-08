package main

import (
	"encoding/json"
	"fmt" // formatter

	"log"       // log errors
	"math/rand" // generate random from math
	"net/http"  // to work with http and api
	"strconv"   // string converter

	"github.com/gorilla/mux" //https://github.com/gorilla/mux
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slick Book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000)) // Mock Id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) //slick index in books
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // Mock Id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete selected book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) //slick index in books
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("Hello World")

	// Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement BD
	books = append(books, Book{ID: "1", Isbn: "2345", Title: "Book one",
		Author: &Author{Firstname: "Krish", Lastname: "Don"}})
	books = append(books, Book{ID: "2", Isbn: "9087", Title: "Book two",
		Author: &Author{Firstname: "John", Lastname: "Will"}})
	books = append(books, Book{ID: "3", Isbn: "4567", Title: "Book three",
		Author: &Author{Firstname: "Cris", Lastname: "Quick"}})

	// Route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	// run the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
