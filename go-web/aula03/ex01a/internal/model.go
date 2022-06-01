package products

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Tipo  string  `json:"tipo"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}
