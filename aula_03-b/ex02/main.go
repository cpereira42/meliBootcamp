package main

import "fmt"

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []produtos
}

type produtos struct {
	Nome       string
	Preco      string
	Quantidade string
}

type Ecommerce struct {
	User []Usuario
	Prod []produtos
}

func main() {
	loja := Ecommerce{}
	loja.Prod = []produtos{}
	loja.User = []Usuario{}

	loja.User = append(loja.User, Usuario{Nome: "cezar", Sobrenome: "Augusto", Email: "cezar@"})

	prod1 := produtos{}
	prod1 = novoProduto("Mac", "5.000,00", "10")
	loja.Prod = append(loja.Prod, prod1)
	loja.Prod = append(loja.Prod, novoProduto("Notebook", "5.000,00", "10"))
	loja.Prod = append(loja.Prod, novoProduto("TV", "3.000,00", "5"))

	adicionaProduto(&loja.User[0], loja.Prod[0], 5)
	adicionaProduto(&loja.User[0], loja.Prod[1], 5)
	adicionaProduto(&loja.User[0], loja.Prod[2], 5)

	listaProdutos(&loja.User[0].Produtos)
	fmt.Println("------Apagando----------")
	deletaProdutos(&loja.User[0])
	fmt.Println("------Apagado----------")
	fmt.Println("------Listando----------")
	listaProdutos(&loja.User[0].Produtos)
}

func deletaProdutos(u *Usuario) {
	u.Produtos = nil
}

func listaProdutos(p *[]produtos) {
	for _, prod := range *p {
		fmt.Println("Nome: ", prod.Nome, "Preco: ", prod.Preco, "Preco: ", prod.Preco)
	}
}

func adicionaProduto(u *Usuario, p produtos, quantidade int) {
	u.Produtos = append(u.Produtos, p)
}

func novoProduto(valores ...string) produtos {
	var t = produtos{}
	t.Nome = valores[0]
	t.Preco = valores[1]
	t.Quantidade = valores[2]
	return t
}
