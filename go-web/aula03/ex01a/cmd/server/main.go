package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meliBootcamp/go-web/aula03/ex01a/cmd/server/handler"
	products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/store"
)

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
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
