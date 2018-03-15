//https://www.youtube.com/watch?v=SonwZ6MF5BE

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
//Book Struct {model}
//Init books var as a slice Book struct
var books []Book

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"Isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}
type  Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json" )
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json" )
	params := mux.Vars(r)
	//loop through books and find the id 
	for _,item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	json.NewEncoder(w).Encode(&Book{})
	//json.NewEncoder(w).Encode(books)
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json" )
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json" )
	params := mux.Vars(r)
	//loop through books and find the id 
	for index,item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index + 1:]...) //delete
			var book Book //new book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book) //add new book
			json.NewEncoder(w).Encode(book)
			return
		}

	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json" )
	params := mux.Vars(r)
	//loop through books and find the id 
	for index,item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID:"1", Isbn:"3232323", Title:"Book one", Author:&Author{Firstname:"Json", Lastname:"Doe"}})
	books = append(books, Book{ID:"2", Isbn:"4444444", Title:"Book two", Author:&Author{Firstname:"steve", Lastname:"Doe"}})
	books = append(books, Book{ID:"3", Isbn:"3232323", Title:"Book three", Author:&Author{Firstname:"Json", Lastname:"Doe"}})


	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",r))
}
