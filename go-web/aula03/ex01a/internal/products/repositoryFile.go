package products

import (
	"fmt"

	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/store"
)

var ps []Product = []Product{}
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, tipo string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, tipo string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].ID, nil
}

func (r *repository) Store(id int, name, tipo string, count int, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, tipo, count, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	//lastID = p.ID
	return p, nil
}

func (r *repository) Update(id int, name, tipo string, count int, price float64) (Product, error) {

	p := Product{Name: name, Tipo: tipo, Count: count, Price: price}
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			p = ps[i]
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) Delete(id int) error {
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			ps = append(ps[:index], ps[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("produto não encontrado")

}
