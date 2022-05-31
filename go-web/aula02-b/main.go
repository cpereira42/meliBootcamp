package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Usuarios []usuario

//var lastID int = 1
var usuarios Usuarios

type usuario struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome" validate:"required"`
	Sobrenome   string  `json:"sobrenome" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	Idade       int     `json:"idade" validate:"required"`
	Altura      float32 `json:altura validate:"required"`
	Ativo       bool    `json:"ativo" validate:"required"`
	DataCriacao string  `json:"datacriacao" validate:"required" `
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

	var lastId int = 1

	jsonData := `[{
		"Nome":"Cezar",
		"SobreNome":"Augusto",
		"Email":"cezar@gmasils",
		"Idade":36,
		"Altura":1.78,
		"Ativo":true,
		"DataCriacao":"55/10/2022"
	}]`

	var u []usuario
	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		log.Fatal(err)
	}
	insertUser(u, &lastId)

	router := gin.Default()
	router.GET("bem-vindo", printName)
	group := router.Group("/usuarios")
	{
		group.GET("getall", getAll())
		group.GET("getid/:id", getId())
		group.GET("getnome/", getNome())
		group.POST("adduser/", addUser(&lastId))
	}

	router.Run()
}

func getNome() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nome := c.Query("nome")
		for _, pd := range usuarios {
			fmt.Println(pd.Nome)
			if pd.Nome == nome {
				c.JSON(200, pd)
				break
			}
		}
	}
	return gin.HandlerFunc(fn)
}

func getId() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, pd := range usuarios {
			if pd.Id == id {
				c.JSON(200, pd)
				return
			}
		}
		c.JSON(404, "Id não encontrado")
	}
	return gin.HandlerFunc(fn)
}

func getAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		for _, pd := range usuarios {
			c.JSON(200, pd)
		}
	}
	return gin.HandlerFunc(fn)
}

func addUser(lastId *int) gin.HandlerFunc {

	var validate *validator.Validate
	validate = validator.New()

	fn := func(c *gin.Context) {
		var user usuario

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//

		err := validate.Struct(user)
		fmt.Println("teste", user.Nome)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
				return
			}
			for _, err := range err.(validator.ValidationErrors) {
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "O campo " + err.Field() + " é obrigatório",
					})
					return
				}

			}
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//

		user.Id = *lastId
		usuarios = append(usuarios, user)
		*lastId++
		c.JSON(200, gin.H{
			"data": user,
		})
		//

	}
	return gin.HandlerFunc(fn)
}

func checkErro(user usuario) error {
	var validate *validator.Validate
	validate = validator.New()

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				return errors.New("O campo " + err.Field() + " é obrigatório do tipo" + err.Param())
			}

		}
	}
	return nil
}

func insertUser(u []usuario, lastId *int) {

	for _, us := range u {
		us.Id = *lastId
		usuarios = append(usuarios, us)
		*lastId++
	}
}
