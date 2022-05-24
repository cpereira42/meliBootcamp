package main

import ("fmt"
		"bufio"
		"strings"
		"io"
		"strconv"
        "os")

func main(){
	opcao :=  "Sim"
	for opcao == "Sim" || opcao == "sim" {
		exibeMenu()
		opcao = getVal("Deseja adicionar mais ? Digite Sim para continuar ")
	}
	fmt.Println("Muito bem aqui está a lista de produtos")
	abreArquivo()
}

func abreArquivo(){
	arquivo, err := os.Open("lista.csv")
	if err != nil {
		fmt.Println("Falha ao abrir")
	} else {
		exibeCabecalho()
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err := leitor.ReadString('\n')
			if err == io.EOF{
				break
			}
			s := strings.Split(linha, ";")
			preco, _ := strconv.ParseFloat(s[1], 32)
			fmt.Printf("%10s %50.2f %50s", s[0], preco ,s[2])
		}
	}
}

func exibeCabecalho(){
	fmt.Printf("%10s %50s %50s\n", "ID", "Preco", "Quantidade")
}

func exibeMenu(){
	id := getVal("Digite o ID")
	preco := getVal("Digite o o preço")
	qtt := getVal("Digite a quantidade")
	escreve(id, preco, qtt)
}

func getVal(msg string) string {
	fmt.Println(msg)
	resp := ""
	fmt.Scanf("%s",&resp)
	return resp
}

func escreve(valores ... string) bool {
	arquivo, err := os.OpenFile("lista.csv",os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
	if err != nil{
        fmt.Println(err)
		fmt.Println("Falha ao gravar")
		return false
    } else {
		arquivo.WriteString(valores[0]+";"+valores[1]+";"+valores[2]+"\n")
		fmt.Println("Gravado com sucesso!!")
		arquivo.Close()
		return true
	}

}
