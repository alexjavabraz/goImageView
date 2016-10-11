package main

import "fmt"

func main() {
	retorno := soma(10, 20)

	fmt.Println(retorno)
}

func soma(n, m int) int {
	return n + m
}
