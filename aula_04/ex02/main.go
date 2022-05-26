package main

import (
	"errors"
	"fmt"
)

func main() {
	salario := 1500

	if salario < 15000 {
		fmt.Println(errors.New("o salário digitado não está dentro do valor mínimo"))
		return
	}
	fmt.Println("Necessário pagamento de imposto")
}
