package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Servico struct {
	Nome   string
	Preco  float64
	minuto int
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func main() {

	prod := []Produto{}
	prod = append(prod, Produto{"Prod1", 100.5, 1})
	prod = append(prod, Produto{"Prod2", 50, 2})
	prod = append(prod, Produto{"Prod3", 20, 4})

	serv := []Servico{}
	serv = append(serv, Servico{"Serv1", 100.5, 10})
	serv = append(serv, Servico{"Serv2", 50, 2})
	serv = append(serv, Servico{"Serv3", 20, 4})

	manu := []Manutencao{}
	manu = append(manu, Manutencao{"manu1", 100.5})
	manu = append(manu, Manutencao{"manu2", 50})
	manu = append(manu, Manutencao{"manu3", 20})

	rprod := make(chan float64)
	rserv := make(chan float64)
	rmanu := make(chan float64)

	go somaProduto(prod, rprod)
	go somaServico(serv, rserv)
	go somarManutencao(manu, rmanu)

	pro := <-rprod + <-rserv + <-rmanu
	fmt.Println("soma =", pro)
}

func somaProduto(p []Produto, out chan<- float64) {
	var sum float64 = 0
	for _, prod := range p {
		sum += float64(prod.Preco) * float64(prod.Quantidade)
	}
	out <- float64(sum)

	close(out)
}

func somaServico(s []Servico, out chan<- float64) {
	total := 0
	for _, serv := range s {
		if serv.minuto < 30 {
			total += 30
		} else {
			total += serv.minuto
		}
		total += serv.minuto * int(serv.Preco) / 60
	}
	out <- float64(total)

	close(out)
}

func somarManutencao(m []Manutencao, out chan<- float64) {
	var total float64 = 0
	for _, manu := range m {
		total += manu.Preco
	}
	out <- float64(total)
	close(out)
}
