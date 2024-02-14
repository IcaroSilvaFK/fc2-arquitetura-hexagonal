package application

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

var (
	ErrorPrice = errors.New("THE PRICE MUST BE GREATER THAN ZERO TO ENABLE THE PRODUCT")
)

func NewProduct(name, status string, price float64) *Product {
	return &Product{
		ID:     uuid.NewString(),
		Name:   name,
		Price:  price,
		Status: status,
	}
}

// func (p *Product) IsValid() (bool, error) {

// }

func (p *Product) Enable() error {

	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return ErrorPrice
}

// func (p *Product) Disable() error
// func (p *Product) GetID() string
// func (p *Product) GetName() string
// func (p *Product) GetStatus() string
// func (p *Product) GetPrice() float64
