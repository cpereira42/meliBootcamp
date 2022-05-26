package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg string
}

func (m *myError) Error() string {
	return fmt.Sprintf(" %v", m.msg)
}

func myCustomError(salario int) (string, error) {
	if salario < 15000 {
		return "", &myError{msg: "erro: o salário digitado não está dentro do valor mínimo"}
	}
	return "Necessário pagamento de imposto", nil
}

func main() {
	var salario int = 50000
	msg, err := myCustomError(salario)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(msg)
}
