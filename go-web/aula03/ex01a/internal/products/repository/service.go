package products

type Service interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Store(name, tipo string, count int, price float64) (Product, error)
	Update(id int, name, tipo string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetId(id int) (Product, error) {
	ps, err := s.repository.GetId(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Store(name, tipo string, count int, price float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++

	product, err := s.repository.Store(lastID, name, tipo, count, price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name, tipo string, count int, price float64) (Product, error) {
	product, err := s.repository.Update(id, name, tipo, count, price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	product, err := s.repository.UpdateName(id, name)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
