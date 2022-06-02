package products

import (
	"fmt"
)

type repositoryRam struct{}

func NewRepositoryRam() Repository {
	return &repositoryRam{}
}

func (r *repositoryRam) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repositoryRam) LastID() (int, error) {
	return lastID, nil
}

func (r *repositoryRam) Store(id int, name, tipo string, count int, price float64) (Product, error) {
	p := Product{id, name, tipo, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repositoryRam) Update(id int, name, tipo string, count int, price float64) (Product, error) {

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

func (r *repositoryRam) UpdateName(id int, name string) (Product, error) {
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

func (r *repositoryRam) Delete(id int) error {
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
