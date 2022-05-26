package main

import "fmt"

func main() {

	salario := 80000

	if salario < 15000 {
		err := fmt.Errorf("erro: o mínimo tributávels é 15.000 e o salário informado é %d", salario)
		fmt.Println(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}

}
