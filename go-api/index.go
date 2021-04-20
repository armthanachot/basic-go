package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type books struct {
	bookName   string
	authorName string
	price      int
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to index")
}
func getAllBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to books")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	allbook := books{
		bookName:   "JAVA",
		authorName: "Thanachot",
		price:      500,
	}
	json.NewEncoder(w).Encode(allbook)
	fmt.Fprint(w, allbook)

}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", getAllBook)
	http.ListenAndServe(":8080", nil)

}

func main() {
	handleRequest()
}
