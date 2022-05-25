package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	opcao := "Sim"
	opcao = getVal("Deseja adicionar ?")
	for opcao == "Sim" || opcao == "sim" {
		exibeMenu()
		opcao = getVal("Deseja adicionar mais ? Digite Sim para continuar ")
	}
	fmt.Println("Muito bem aqui está a lista de produtos")
	fmt.Printf("Total %2.f\n", abreArquivo())
}

func abreArquivo() float32 {
	var total float32 = 0.
	arquivo, err := os.Open("lista.csv")
	if err != nil {
		fmt.Println("Falha ao abrir")
	} else {
		exibeCabecalho()
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.Trim(linha, "\n")
			if err == io.EOF {
				break
			}
			s := strings.Split(linha, ";")
			preco, _ := strconv.ParseFloat(s[1], 32)

			qtt, _ := strconv.ParseFloat(s[2], 32)
			total += float32(preco * qtt)
			fmt.Printf("%10s %18.2f %20.2f\n", s[0], preco, qtt)
		}
	}
	fmt.Println("Total ", total)
	return total
}

func exibeCabecalho() {
	fmt.Printf("%10s %18s %20s\n", "ID", "Preco", "Quantidade")
}

func exibeMenu() {
	id := getVal("Digite o ID")
	preco := getVal("Digite o o preço")
	qtt := getVal("Digite a quantidade")
	escreve(id, preco, qtt)
}

func getVal(msg string) string {
	fmt.Println(msg)
	resp := ""
	fmt.Scanf("%s", &resp)
	return resp
}

func escreve(valores ...string) bool {
	arquivo, err := os.OpenFile("lista.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Falha ao gravar")
		return false
	} else {
		arquivo.WriteString(valores[0] + ";" + valores[1] + ";" + valores[2] + "\n")
		fmt.Println("Gravado com sucesso!!")
		arquivo.Close()
		return true
	}

}
