package products_test

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

func Test_Repository(t *testing.T) {

	var prodNew = products.Product{ID: 3, Name: "Toshiba", Tipo: "TV", Count: 6, Price: 7000}
	var prod = products.Product{ID: 1, Name: "Toshiba", Tipo: "TV", Count: 6, Price: 7000}
	var prod2 = products.Product{ID: 2, Name: "Samsung", Tipo: "TV", Count: 6, Price: 7000}

	produtos := []products.Product{prod, prod2}

	t.Run("test Get All", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(produtos, nil)

		service := products.NewService(repo)
		ps, err := service.GetAll()

		assert.Nil(t, err)
		assert.True(t, len(ps) > 0)
		assert.Equal(t, produtos[0], ps[0])
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(nil, fmt.Errorf("Falha ao buscar produtos"))

		service := products.NewService(repo)
		_, err := service.GetAll()
		assert.Equal(t, fmt.Errorf("Falha ao buscar produtos"), err)
	})

	t.Run("TestDeleteOk", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Delete", 2).Return(nil)

		service := products.NewService(repo)
		err := service.Delete(2)
		assert.Nil(t, err)
	})

	t.Run("TestDeleteFail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Delete", 9).Return(fmt.Errorf("produto não encontrado"))

		service := products.NewService(repo)
		err := service.Delete(9)

		assert.Equal(t, fmt.Errorf("produto não encontrado"), err)
	})

	t.Run("TestStoreOk", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("LastID").Return(2, nil)
		repo.On("Store", 3, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(prodNew, nil)

		service := products.NewService(repo)
		ps, err := service.Store(prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

		assert.Nil(t, err)
		assert.Equal(t, prodNew, ps)
	})

	t.Run("TestStoreFail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("LastID").Return(2, nil)
		repo.On("Store", 3, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(products.Product{}, fmt.Errorf("Falha ao gravar produtos"))

		service := products.NewService(repo)
		_, err := service.Store(prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

		assert.Equal(t, fmt.Errorf("Falha ao gravar produtos"), err)
	})

	t.Run("TestStoreFailStore", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("LastID").Return(0, fmt.Errorf("Falha ao gravar produtos"))
		repo.On("Store", 3, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(products.Product{}, fmt.Errorf("Falha ao gravar produtos"))

		service := products.NewService(repo)
		_, err := service.Store(prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

		assert.Equal(t, fmt.Errorf("Falha ao gravar produtos"), err)
	})

	t.Run("TestUpdateOk", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Update", 2, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(prod, nil)

		service := products.NewService(repo)
		ps, err := service.Update(2, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

		assert.Nil(t, err)
		assert.Equal(t, prod.Name, ps.Name)
	})

	t.Run("TestUpdateFail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Update", 9, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(products.Product{}, fmt.Errorf("Falha ao atualizar produtos"))

		service := products.NewService(repo)
		_, err := service.Update(9, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

		assert.Equal(t, fmt.Errorf("Falha ao atualizar produtos"), err)
	})

	t.Run("TestUpdateNameOk", func(t *testing.T) {
		var prod3 = products.Product{ID: 2, Name: "teste", Tipo: "TV", Count: 6, Price: 7000}
		repo := &mocks.Repository{}
		repo.On("UpdateName", 2, prod3.Name).Return(prod3, nil)

		service := products.NewService(repo)
		ps, err := service.UpdateName(2, "teste")

		assert.Nil(t, err)
		assert.Equal(t, prod3.Name, ps.Name)
	})

	t.Run("TestUpdateNameFail", func(t *testing.T) {
		var prod3 = products.Product{ID: 2, Name: "teste", Tipo: "TV", Count: 6, Price: 7000}
		repo := &mocks.Repository{}
		repo.On("UpdateName", 9, prod3.Name).Return(products.Product{}, fmt.Errorf("Falha ao atualizar produtos"))

		service := products.NewService(repo)
		_, err := service.UpdateName(9, "teste")

		assert.Equal(t, fmt.Errorf("Falha ao atualizar produtos"), err)
	})

}

/*func TestUpdateName(t *testing.T) {
	repo := &mocks.Repository{}
	var prodNew = products.Product{ID: 3, Name: "Toshiba", Tipo: "TV", Count: 6, Price: 7000}
	repo.On("LastID").Return(2, nil)
	repo.On("Store", 3, prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price).Return(prodNew, nil)

	service := products.NewService(repo)
	ps, err := service.Store(prodNew.Name, prodNew.Tipo, prodNew.Count, prodNew.Price)

	assert.Nil(t, err)
	assert.Equal(t, prodNew, ps)
	//assert.Equal(t, "1", "1")
}*/

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
