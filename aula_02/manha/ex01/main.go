package main

import "fmt"

const(
	ImpostoMedio = 17
	ImpostoAuto = 10
)



func main(){
	fmt.Println("Imposto retido", calcImposto(50000), "para o salario de 50.000")
	fmt.Println("Imposto retido", calcImposto(150000), "para o salario de 150.000")
}

func calcImposto(salario float32) float32 {

	if salario == 50000{
		return salario * ImpostoMedio / 100
	}
	if salario == 150000 {
		return salario * ImpostoAuto / 100
	}
	return 0
}