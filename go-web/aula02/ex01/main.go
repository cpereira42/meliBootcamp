package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuarios []usuario

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

	router := gin.Default()
	router.GET("bem-vindo", printName)

	var u []usuario
	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		log.Fatal(err)
	}

	group := router.Group("/usuarios")
	{
		group.GET("getall", getAll(u))
		group.GET("getid/:id", getId(u))
		group.GET("getnome/", getNome(u))
	}

	router.Run()
}

func getNome(u []usuario) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nome := c.Query("nome")
		for _, pd := range u {
			fmt.Println(pd.Nome)
			if pd.Nome == nome {
				c.JSON(200, pd)
				break
			}
		}
	}
	return gin.HandlerFunc(fn)
}

func getId(u []usuario) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, pd := range u {
			if pd.Id == id {
				c.JSON(200, pd)
				return
			}
		}
		c.JSON(404, "Id não encontrado")
	}
	return gin.HandlerFunc(fn)
}

func getAll(u []usuario) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		for _, pd := range u {
			c.JSON(200, pd)
		}
	}
	return gin.HandlerFunc(fn)
}
