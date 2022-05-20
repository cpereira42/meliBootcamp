package main

import "fmt"

func main() {
	var idade, periodo int
	var salario float32
	

	for {
		fmt.Println("Digite a idade")
		fmt.Scan(&idade)

		if idade <= 22 {
			fmt.Println("Não tem idade mínima para fazer empréstimo\n")
			if questiona() {
				continue
			} else {
				break
			}
		}

		fmt.Println("Digite o tempo trabalhado em meses")
		fmt.Scan(&periodo)
		if periodo <= 12 {
			fmt.Println("Não tem contribuição mínima\n")
			if questiona() {
				continue
			} else {
				break
			}
		}

		fmt.Println("Digite o salário")
		fmt.Scan(&salario)
		if salario <= 100000 {
			fmt.Println("Emprestimo com juros")
		} else {
			fmt.Println("Emprestimo sem juros")
		}

		if questiona() {
			continue
		} else {
			break
		}

	}
}

func questiona() bool{

	var continua string
	fmt.Println("\nDigite 1, para continuar")
	fmt.Scan(&continua)
	if (continua == "1"){
		return true
	} else {
		return false
	}


	
}

