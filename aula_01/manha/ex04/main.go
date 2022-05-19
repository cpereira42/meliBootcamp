package main

import ("fmt")

func main(){
	/* var sobrenome string = "Silva" ideal colocar cameCase
	var idade int = "25" - Errado, remover aspas duplas
	boolean := "false"; - Declaração invalida
	var salario string = 4585.90 Tipo invalido deveria ser float32
	var nome string = "Fellipe" */

	var sobreNome string = "Silva"
	var idade int = 25
	boolean := false;
	var salario float32 = 4585.90
	var nome string = "Fellipe"

	fmt.Println("Nome", nome )
	fmt.Println("Idade", idade )
	fmt.Println("Sobrenome", sobreNome)
	fmt.Println("teste", boolean)
	fmt.Println("Salario", salario)




}