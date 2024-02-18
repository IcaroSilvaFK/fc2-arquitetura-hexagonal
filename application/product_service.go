package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(p ProductPersistenceInterface) *ProductService {

	return &ProductService{
		Persistence: p,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {

	p := NewProduct(name, price)

	_, err := p.IsValid()

	if err != nil {
		return nil, err
	}

	r, err := s.Persistence.Save(p)

	if err != nil {
		return nil, err
	}

	return r, nil

}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {

	err := product.Enable()

	if err != nil {
		return nil, err
	}

	r, err := s.save(product)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return nil, err
	}

	r, err := s.save(product)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ProductService) save(p ProductInterface) (ProductInterface, error) {

	p, err := s.Persistence.Save(p)

	if err != nil {
		return nil, err
	}
	return p, nil
}
