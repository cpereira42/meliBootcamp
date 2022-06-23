package products

import (
	"fmt"

	"github.com/meliBootcamp/go-web/aula03/ex01a/pkg/store"
)

type repositoryDisc struct {
	db store.Store
}

func NewRepositoryDisc(db store.Store) Repository {
	return &repositoryDisc{
		db: db,
	}
}

func (r *repositoryDisc) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repositoryDisc) GetId(id int) (Product, error) {
	var ps Product
	return ps, nil
}

func (r *repositoryDisc) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].ID, nil
}

func (r *repositoryDisc) Store(id int, name, tipo string, count int, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, tipo, count, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repositoryDisc) Update(id int, name, tipo string, count int, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{Name: name, Tipo: tipo, Count: count, Price: price}
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repositoryDisc) UpdateName(id int, name string) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			p := ps[i]
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repositoryDisc) Delete(id int) error {
	var index int
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].ID == id {
			index = i
			ps = append(ps[:index], ps[index+1:]...)
			if err := r.db.Write(ps); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("produto não encontrado")

}
