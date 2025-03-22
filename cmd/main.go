package main

import "fmt"

func main() {
	fmt.Println("Qual seu nome? ")
	var nome string
	fmt.Scan(&nome)
	fmt.Println("Ola, ", nome)
}
