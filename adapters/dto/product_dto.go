package dto

import (
	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
)

type ProductDTO struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func (p *ProductDTO) Bind(product *application.Product) (*application.Product, error) {

	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()

	if err != nil {
		return nil, err
	}

	return product, nil
}
