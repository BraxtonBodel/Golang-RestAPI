package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct (model)
type Author struct {
	Firstname string `json:"firstname"`
	Lasttname string `json:"lastname"`
}

//Init books var as a slice Book Struct
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r) //Get params
	//loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create book
func createBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) //MocK ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//init router
	r := mux.NewRouter()

	//Mock Data - @todo - implementation DB
	books = append(books, Book{ID: "1", Isbn: "223212", Title: "100 años de soledad", Author: &Author{
		Firstname: "Grabiel", Lasttname: "Garcia Marquez"}})

	books = append(books, Book{ID: "2", Isbn: "223213", Title: "El quijote", Author: &Author{
		Firstname: "Grabiel", Lasttname: "Garcia Marquez"}})

	books = append(books, Book{ID: "3", Isbn: "223214", Title: "El mio cid", Author: &Author{
		Firstname: "Grabiel", Lasttname: "Garcia Marquez"}})

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
