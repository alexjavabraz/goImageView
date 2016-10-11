package main

import "fmt"

func main() {
	fmt.Printf("%s tem %d anos. \n", "Fernando", 29)
	imprimirDados("Fernando", 29)
}

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos. \n", nome, idade)
}
