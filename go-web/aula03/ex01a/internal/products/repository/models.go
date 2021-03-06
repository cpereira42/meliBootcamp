package products

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Tipo  string  `json:"tipo"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

type Repository interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Store(id int, name, tipo string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, tipo string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}
