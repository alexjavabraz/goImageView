package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (h Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello Struct!")
}

func main() {
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
	http.Handle("localhost:4000/string", String("é uma string"))
	http.Handle("localhost:4000/struct", &Struct{"Hello", ":", "Gophers!"})
}
