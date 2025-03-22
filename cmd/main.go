package main

import "fmt"

func sayHello(name string) string {
	r := fmt.Sprintf("Ola, %s", name)
	return r
}

func main() {
	fmt.Println("Qual seu nome? ")
	var nome string
	fmt.Scan(&nome)
	fmt.Println(sayHello(nome))
}
