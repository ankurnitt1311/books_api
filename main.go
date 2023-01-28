package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
)

//Get All Books function
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(books)
}

//get single book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//updateBook
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)
}

//deleteBook
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Book struct(model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Autor struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	//Init Router
	r := mux.NewRouter()
	books = append(books, Book{
		ID:     "1",
		Isbn:   "12345",
		Title:  "First Book",
		Author: &Author{Firstname: "Anku", Lastname: "Kulhari"},
	})
	books = append(books, Book{
		ID:     "2",
		Isbn:   "12346",
		Title:  "second Book",
		Author: &Author{Firstname: "Ankur", Lastname: "Choudhary"},
	})
	//Route Handler /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Server is up and running")
	log.Fatal(http.ListenAndServe(":8080", r))

}
