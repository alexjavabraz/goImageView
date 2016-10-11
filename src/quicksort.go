package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade> exemplo go run conversor.go 10 11 12 quilometros")
		os.Exit(1)
	}
}
