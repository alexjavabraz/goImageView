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

func HelloServer(w http.ResponseWriter, req *http.Request) {
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