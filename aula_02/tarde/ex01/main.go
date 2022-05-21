package main

import "fmt"

type Pessoa struct {
	Nome			string
	Sobrenome		string
	Rg				string
	DataAdmissao	string
}

func main (){
	aluno := Pessoa{}
	fmt.Println("Digite o nome")
	fmt.Scan(&aluno.Nome)
	fmt.Println("Sobrenome")
	fmt.Scan(&aluno.Sobrenome)
	fmt.Println("Digite o RG")
	fmt.Scan(&aluno.Rg)
	fmt.Println("Digite a Data de Admissão")
	fmt.Scan(&aluno.DataAdmissao)
	aluno.imprimir()
}

func (a Pessoa) imprimir(){
	fmt.Printf("Nome %s %s \nRG : %s \nData Admissão %s\n",a.Nome,a.Sobrenome,a.Rg,a.DataAdmissao)
}
