package main

import (
	"fmt"
	"strconv"
)

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
}

func main() {

	pessoa := Usuario{"Cezar", "Augusto", 36, "cezaracp1@gmail.com", "123456"}

	println(pessoa.Nome, pessoa.Sobrenome, pessoa.Idade, pessoa.Email, pessoa.Senha)
	mudarNome(&pessoa)
	mudarIdade(&pessoa)
	mudarEmail(&pessoa)
	mudarSenha(&pessoa)
	println(pessoa.Nome, pessoa.Sobrenome, pessoa.Idade, pessoa.Email, pessoa.Senha)
	pessoa.mudarNome2()
	pessoa.mudarIdade2()
	pessoa.mudarEmail2()
	pessoa.mudarSenha2()
	println(pessoa.Nome, pessoa.Sobrenome, pessoa.Idade, pessoa.Email, pessoa.Senha)
}

func mudarNome(u *Usuario) {
	u.Nome = getVal("Alterando -> " + u.Nome + ", Digite o novo nome")
	u.Sobrenome = getVal("Alterando sobrenome ->" + u.Sobrenome + ", Digite o novo sobrenome")
}

func mudarIdade(u *Usuario) {
	ret, _ := strconv.Atoi(getVal("Digite a nova idade"))
	u.Idade = ret
}

func mudarEmail(u *Usuario) {
	u.Email = getVal("Alterando ->" + u.Email + ", Digite o novo email")
}

func mudarSenha(u *Usuario) {
	u.Senha = getVal("Alterando Senha, Digite a nova senha")
}

func getVal(msg string) string {
	fmt.Println(msg)
	resp := ""
	fmt.Scanf("%s", &resp)
	return resp
}

func (u *Usuario) mudarSenha2() {
	u.Senha = getVal("Alterando Senha, Digite a nova senha")
}

func (u *Usuario) mudarNome2() {
	u.Nome = getVal("Alterando -> " + u.Nome + ", Digite o novo nome")
	u.Sobrenome = getVal("Alterando sobrenome ->" + u.Sobrenome + ", Digite o novo sobrenome")
}

func (u *Usuario) mudarIdade2() {
	ret, _ := strconv.Atoi(getVal("Digite a nova idade"))
	u.Idade = ret
}

func (u *Usuario) mudarEmail2() {
	u.Email = getVal("Alterando ->" + u.Email + ", Digite o novo email")
}
