package main

import "fmt"

const(
	CatA = 3000
	CatB = 1500
	CatC = 1000
	ExtraA = 50
	ExtraB = 20
	HoraExtra = 160
)

func main(){
	fmt.Println("O Salario da Categoria C é",calcSalario(162,"C"))
	fmt.Println("O Salario da Categoria B é",calcSalario(176,"B"))
	fmt.Println("O Salario da Categoria A é",calcSalario(172,"A"))

}

func calcSalario(horas int, categoria string) int{

	switch categoria{
	case "A","a":
		if (horas <= HoraExtra){
			return horas * CatA
		}else {
			return (horas * CatA) + (horas * CatA) * ExtraA / 100
		}
	case "B", "b":
		if (horas <= HoraExtra){
			return horas * CatB
		}else {
			return (horas * CatB) + (horas * CatB) * ExtraB / 100
		}
	case "C", "c":
		return horas * CatC

	default:
		return 0
	}
}


