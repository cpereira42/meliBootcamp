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
	prod1 :=novoProduto(Pequeno,"notebook",1500)
	prod2 :=novoProduto(Medio,"Ar Condicionado",9000)
	prod3 :=novoProduto(Grande,"Geladeira",7500)
	loja1 := novaLoja()
	loja1.Adicionar(prod1)
	loja1.Adicionar(prod2)
	loja1.Adicionar(prod3)
	total, qtt := loja1.Total()
	fmt.Printf("A Loja 1 tem %d produtos totalizando %.2f",qtt, total)
}

func novaLoja() loja{
	return loja{}
}

func (l *loja) Adicionar(prod produto){
	l.ListaProdutos = append(l.ListaProdutos,prod)
} 

func (l loja) Total() (float32,int) {

	var sum float32 = 0
	qtt :=0
	for _, produto := range l.ListaProdutos{
		sum += float32(produto.calcularCusto() + produto.Preco)
		qtt++
	}
	return sum, qtt
}

func (a produto) calcularCusto() float32{
	
	switch a.Tipo {
	case Pequeno:
		return 0
	case Medio:
		return a.Preco * 3 / 100
	case Grande:
		return a.Preco * 6 / 100 + 2500
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


