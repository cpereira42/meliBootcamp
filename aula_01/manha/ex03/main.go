package main

import ("fmt")

func main(){

	/* var 1nome string //Errado não pode comecar com numero
	var int idade // Errado, ordem invertida
	1sobrenome :=6 // Errado - deveria estar entre aspas duplas e não deveriqa começar com numero
	var licenca_para_dirigir = true //Errado - O formato correto seria em camelCase ex licencaParaDirigir
	var estatura da pessoa int // Errado, não pode conter espaço estaturaDaPessoa
	quantidadeDeFilhos :=2 // está correto */

	var nome string
	var idade int = 18
	sobreNome := "6"
	var licencaParaDirigir = true 
	var estaturaDaPessoa int 
	quantidadeDeFilhos :=2 
	

	fmt.Println("Nome", nome )
	fmt.Println("Idade", idade )
	fmt.Println("Sobrenome", sobreNome)
	fmt.Println("Licença para Dirigir", licencaParaDirigir)
	fmt.Println("Estatura da Pessoa", estaturaDaPessoa)
	fmt.Println("Quantidade de filhos", quantidadeDeFilhos)




}