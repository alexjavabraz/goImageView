package main

import (
	"fmt"
	"os"
	"strconv"
)

func soma(n, m int) int {
	return n + m
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade> exemplo go run conversor.go 10 11 12 quilometros")
		os.Exit(1)
	}

	soma(1, 2)

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		fmt.Printf("Valor %s na posição %d \n", v, i)

		valorOrigem, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf(" O valor %s na posição %d não é um número válido! \n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s \n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)

	}

}
