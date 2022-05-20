package main

import "fmt"
import "os"

func main(){

	var nome string
	var opcao int

	
	if len(os.Args) > 2{
		fmt.Println("Digite apenas uma palavra")
	} else if len(os.Args) == 2 {
		nome =os.Args[1]
	} else{
		fmt.Println("Digite a palavra")
		fmt.Scanf("%s", &nome)
	}
	
	fmt.Println("A Palavra", nome, "tem",len(nome),"letras")
	
	fmt.Println("Qual estilo você quer soletrar :")
	fmt.Println("1 - Standard")
	fmt.Println("2 - RangeFor")
	fmt.Println("3 - Loop Infinito")
	fmt.Println("4 - For while")
	fmt.Scan(&opcao)

	switch opcao {
	case 1 :
		soletrarStandard(nome)
	case 2:
		soletrarRangeFor(nome)
	case 3:
		soletrarLoopInfinito(nome)
	case 4:
		soletraForWhile(nome)
	default:
		fmt.Println("Escolha inválida")
	}

}

func soletrarStandard(nome string){
	fmt.Println("Vamos soletrar Standard for")
	for i := 0; i < len(nome); i++ {
		fmt.Println(string(nome[i]))
	}
}

func soletrarRangeFor(nome string){
	fmt.Println("Vamos soletrar Range for")
	for i :=range nome{
		fmt.Println(string(nome[i]))
	}
}

func soletrarLoopInfinito(nome string){
	fmt.Println("Vamos soletrar For infinito")
	i:=0
	for {
		if i >= len(nome) {
			break
		} else {
			fmt.Println(string(nome[i]))
		}
		i++
	}
}

func soletraForWhile(nome string){
	fmt.Println("Vamos soletrar For While")
	i := 0
	for i < len(nome){
		fmt.Println(string(nome[i]))
		i++
	}
}