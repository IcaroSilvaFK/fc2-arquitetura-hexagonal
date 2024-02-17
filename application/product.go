package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Product struct {
	ID     string  `json:"id" validate:"required,uuid"`
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required,gt=0"`
	Status string  `json:"status,omitempty"`
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

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductWriter
	ProductReader
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

var (
	ErrorPriceIsLessThanZero    = errors.New("THE PRICE MUST BE GREATER THAN ZERO TO ENABLE THE PRODUCT")
	ErrorPriceIsGreaterThanZero = errors.New("THE PRICE MUST BE LESS THAN ZERO TO DISABLE THE PRODUCT")
)

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func (p *Product) IsValid() (bool, error) {

	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("STATUS MUST BE ENABLED OR DISABLED")
	}

	if p.Price < 0 {
		return false, ErrorPriceIsLessThanZero
	}

	if p.Name == "" {
		return false, errors.New("NAME IS REQUIRED")
	}

	validator := validator.New()

	if err := validator.Struct(p); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {

	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return ErrorPriceIsLessThanZero
}

func (p *Product) Disable() error {

	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return ErrorPriceIsGreaterThanZero
}

func (p *Product) GetID() string {
	return p.ID
}
func (p *Product) GetName() string {
	return p.Name
}
func (p *Product) GetStatus() string {
	return p.Status
}
func (p *Product) GetPrice() float64 {
	return p.Price
}
