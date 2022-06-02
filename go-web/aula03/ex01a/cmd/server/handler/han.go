package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/web"
)

type request struct {
	Name  string  `json:"name"`
	Tipo  string  `json:"tipo"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

type requestName struct {
	Name string `json:"name"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} request
// @Failure 400 {object} web.Response "We need ID!!"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID!!"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		p, err := c.service.Store(req.Name, req.Tipo, req.Count, req.Price)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id  path string true "id"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Responsec
// @Failure 400 {object} web.Response "We need ID!!"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{id} [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "O nome do produto é obrigatório"))
			return
		}

		if req.Tipo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "O tipo do produto é obrigatório"))
			return
		}

		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "A Quantidade do produto é obrigatória"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "O preço do produto é obrigatório"))
			return
		}
		p, err := c.service.Update(int(id), req.Name, req.Tipo, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id  path string true "id"
// @Param product body requestName true "Product to updateName"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID!!"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{id} [patch]
func (c *Product) UpdateName() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
			return
		}
		var req requestName
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(401, web.NewResponse(400, nil, "O nome do produto é obrigatório"))
			return
		}
		p, err := c.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id  path string true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID!!"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{id} [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("O produto %d foi removido", id), ""))
	}
}
