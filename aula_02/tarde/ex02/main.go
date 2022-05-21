package main

import "fmt"

const(
	Pequeno = "Pequeno"
	Medio = "Medio"
	Grande = "Grande"
)

type Produto interface{
	calcularCusto() float32
}

type Ecommerce interface{
	Total() float32
	Adicionar(produto)
}

type loja struct {
	ListaProdutos	[]produto
}

type produto struct {
	Tipo			string
	Nome			string
	Preco			float32
}

func main(){
	prod1 :=novoProduto(Medio,"notebook",5000)
	prod2 :=novoProduto(Medio,"TV",5000)
	loja1 := novaLoja()
	loja1.Adicionar(prod1)
	loja1.Adicionar(prod2)
	fmt.Println("Total",loja1.Total())
}

func novaLoja() loja{
	return loja{}
}

func (l *loja) Adicionar(prod produto){
	l.ListaProdutos = append(l.ListaProdutos,prod)
} 

func (l loja) Total() float32 {

	var sum float32 = 0
	for _, produto := range l.ListaProdutos{
		sum += float32(produto.calcularCusto() + produto.Preco)
	}
	return sum
}

func (a produto) calcularCusto() float32{
	
	switch a.Tipo {
	case Pequeno:
		return 0
	case Medio:
		return a.Preco * 3 / 100
	case Grande:
		return a.Preco * 5 / 100 + 2500
	default:
		return 0
	}
}

func novoProduto(tipo string, nome string, preco float32) produto{
	return produto{
		Nome : nome,
		Tipo: tipo,
		Preco: preco,
	}
}


