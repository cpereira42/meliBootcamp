package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float32 `json:altura`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"datacriacao"`
}

func printName(c *gin.Context) {
	c.JSON(200, gin.H{
		"messagem": "cezar",
	})
}

func getName() string {
	var u usuario
	arquivo, err := ioutil.ReadFile("usuarios.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(arquivo), &u)
	return u.Nome
}

func main() {

	router := gin.Default()
	router.GET("bem-vindo", printName)

	jsonData := `[{
					"Id":5,
					"Nome":"Cezar",
					"SobreNome":"Augusto",
					"Email":"cezar@gmasils",
					"Idade":36,
					"Altura":1.78,
					"Ativo":true,
					"DataCriacao":"55/10/2022"
				},
				{
					"Nome":"Angélica",
					"Id":6,
					"SobreNome":"Miranda"}
		]`

	var u []usuario
	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		log.Fatal(err)
	}
	router.GET("GetAll", getAll(u))

	router.Run()
}

func getAll(u []usuario) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		for _, pd := range u {
			c.JSON(200, pd)
		}
	}
	return gin.HandlerFunc(fn)
}
