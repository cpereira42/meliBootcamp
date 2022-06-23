package repository_test

import (
	"fmt"
	"testing"

	"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository/mocks"
	"github.com/stretchr/testify/assert"

	products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
)

/*
func createServer() *gin.Engine {

	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../../../products.json")

	repo := products.NewRepositoryDisc(db)

	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProduct_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	objRes := struct {
		Code string
		Data []products.Product
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)

}

func Test_SaveProduct_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/", `{"name":"Tester","tipo":"funcional","count":10, "price":50.00}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}*/

func TestGetAll(t *testing.T) {
	var prod = products.Product{ID: 1, Name: "Toshiba", Tipo: "TV", Count: 6, Price: 7000}
	var prod2 = products.Product{ID: 2, Name: "Samsung", Tipo: "TV", Count: 6, Price: 7000}

	produtos := []products.Product{prod, prod2}

	repo := &mocks.Repository{}
	repo.On("GetAll").Return(produtos, nil)

	service := products.NewService(repo)
	ps, err := service.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(ps) > 0)
	assert.Equal(t, produtos[0], ps[0])
}

func TestDeleteInvalidId(t *testing.T) {
	repo := &mocks.Repository{}
	repo.On("Delete", 9).Return(fmt.Errorf("produto não encontrado"))

	service := products.NewService(repo)
	err := service.Delete(9)

	assert.Equal(t, fmt.Errorf("produto não encontrado"), err)
}

/*
func TestUpdateName(t *testing.T) {

	var prod []Product
	db := store.FileStore{FileName: "../../../db_teste.json", Called: false}
	teste := store.New(store.FileType, "../../../db_teste.json", false)

	repo := NewRepositoryDisc(teste)
	repo.UpdateName(1, "João")
	//teste := db.Called
	db.Read(&prod)

	fmt.Printf(prod[0].Name)

	assert.True(t, db.Called)

	//db := store.New(store.FileType, "../../products.json", false)
	//db := store.

	//st := store.FileStore(store.FileType)

	//db := st.Read("./../products.json")

	/*var prod []Product
	db := store.New(store.FileType, "../../products.json", false)
	db.Read(&prod)
	prod[0].Name = "Joao"

	repo := NewRepositoryDisc(db)

	result, _ := repo.UpdateName(1, "Joao")

		assert.Equal(t, prod[0], result)

}*/
