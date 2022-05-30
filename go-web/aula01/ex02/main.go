package main

import (
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

func main() {

	router := gin.Default()
	router.GET("bem-vindo", printName)
	router.Run()
}
