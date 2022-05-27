package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type usuario struct {
	arquivo  string
	nome     string
	rg       string
	numTel   string
	endereco string
}

func main() {

	user1 := usuario{"0", "cezar", "33", "1", "endereco"}
	id, err := geraId()
	user1.arquivo = id
	if err != nil {
		panic(err)
	}
	fmt.Println("id = ", id)
	lerArquivo(user1, "customers.txt")
	lerArquivo(user1, "customers2.txt")
	lerArquivo(user1, "customers2.txt")
	fmt.Println("Finalizando")

}

func geraId() (string, error) {

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(5)
	if id > 3 {
		fmt.Println("id = ", id)
		return "", errors.New("falha ao gerar arquivo")
	}
	return strconv.Itoa(id), nil
}

func lerArquivo(user usuario, path string) {
	u := usuario{}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	arquivo, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		//log.Println("o arquivo indicado não foi encontrado ou está danificado")
		panic("o arquivo indicado não foi encontrado ou está danificado")
	} else {
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.Trim(linha, "\n")
			if err == io.EOF {
				break
			}
			s := strings.Split(linha, ";")
			u.arquivo = s[0]
			u.nome = s[1]
			u.rg = s[2]
			u.numTel = s[3]
			u.endereco = s[4]
			if user == u {
				panic("Usuario já cadastrado")
			}
		}
		fmt.Println("Cadastrando\n", u, user)
		arquivo.WriteString(user.arquivo + ";" + user.nome + ";" + user.rg + ";" + user.numTel + ";" + user.endereco + ";\n")
	}
	arquivo.Close()
}
