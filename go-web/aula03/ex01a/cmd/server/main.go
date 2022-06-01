package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meliBootcamp/go-web/aula03/ex01a/cmd/server/handler"
	"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products"

	//"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository/disc"

	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/store"
)

func main() {

	/*op := 1

	if op == 1 {


		db := store.New(store.FileType, "../../products.json")
		repo := disc.NewRepository(db)
	} else {
		repo := ram.NewRepository()
	}*/

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Não foi possível abrir o log")
	}

	db := store.New(store.FileType, "../../products.json")
	repo := products.NewRepository(db)
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
