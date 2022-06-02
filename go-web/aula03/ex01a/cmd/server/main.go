package main

import (
	"fmt"
	"log"
	"os"

	docs "github.com/meliBootcamp/go-web/aula03/ex01a/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meliBootcamp/go-web/aula03/ex01a/cmd/server/handler"
	products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/store"
	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/web"
)

//@title Meli Bootcamp API
//@version 1.0
//@description This API Handle MELI Products
//@ termsOfService https://developers.mercadolivre.com.br/pt_br/termos-e-condicoes

//@contact.name API Support
//@contact.url https://developers.mercadolivre.com.br/suporte
//@license.name Apache 2.0
//@licence.url http://www.apache.org/licenses/LICENCE-2.0.html
func main() {

	repo := products.NewRepositoryRam()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Não foi possível abrir o log")
	}

	if len(os.Args) > 1 && os.Args[1] == "Disc" {
		fmt.Println("Vamos gravar no arquivo")
		db := store.New(store.FileType, "../../products.json")
		repo = products.NewRepositoryDisc(db)
	} else {
		fmt.Println("Vamos gravar na memória")
	}

	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := r.Group("/products")
	pr.Use(TokenAuthMiddleware())
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Insira o Token na variável de ambiente")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Middeleware token é necessário"))
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Middeleware token inválido"))
			return
		}
		c.Next()
	}
}
