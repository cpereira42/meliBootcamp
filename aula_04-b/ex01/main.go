package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Tentando ler arquivo!!")
	_, err := os.Open("customers.txt")

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	fmt.Printf("Arquivo lindo com sucesso!!")
}
