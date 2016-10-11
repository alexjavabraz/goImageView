package main

import (
	"fmt"
	"log"
	"net/http" 
)

type Struct struct{


}

func (h Struct) ServeHTTP(
	w http.ResponseWriter, 
	r *http.Request){
	fmt.Fprint(w, "Struct!")
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func HelloServer(w http.ResponseWriter, req *http.Request) Error {

	if(w == nil)
		return &MyError{
						time.Now(),
						"Erro"
					}

	fmt.Fprint(w, "Struct 2 !")
}

func main(){
	var h Struct
	err := http.ListenAndServe("localhost:4000", h)

	http.HandleFunc("/hello", HelloServer)

	if(err != nil){
		log.Fatal(err)
	}
}