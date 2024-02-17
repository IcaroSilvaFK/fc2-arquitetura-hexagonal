package application_test

import (
	"testing"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

func TestShouldEnableProductWhenPriceIsGTOnZero(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	err := product.Enable()

	assert.Nil(t, err)
}

func TestShouldNotEnableProductWhenPriceIsLEOnZero(t *testing.T) {
	product := application.NewProduct("Product 1", 0)

	err := product.Enable()

	assert.NotNil(t, err, err.Error())
	assert.Equal(t, err, application.ErrorPriceIsLessThanZero)
}

func TestShouldNotDisableProductWhenPriceIsGreaterThanZero(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	err := product.Disable()

	assert.NotNil(t, err)
	assert.Equal(t, err, application.ErrorPriceIsGreaterThanZero)
	assert.Equal(t, product.Status, "")
}

func TestShouldDisableProductWhenPriceIsLEZero(t *testing.T) {

	product := application.NewProduct("Product 1", 0)

	err := product.Disable()

	assert.Nil(t, err)
	assert.Equal(t, product.Status, application.DISABLED)
}

func TestShouldValidateProductWhenIsValid(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	isValid, err := product.IsValid()

	assert.Nil(t, err)
	assert.True(t, isValid)
}

func TestShouldValidateProductWhenIsInvalid(t *testing.T) {

	product := application.NewProduct("", 0)

	isValid, err := product.IsValid()

	assert.NotNil(t, err)
	assert.False(t, isValid)
}

func TestShouldValidateProductWhenStatusIsInvalid(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	product.Status = "invalid"

	isValid, err := product.IsValid()

	assert.NotNil(t, err)
	assert.False(t, isValid)
}

func TestShouldValidateProductWhenValueIsInvalid(t *testing.T) {

	product := application.NewProduct("Product 1", -1)

	isValid, err := product.IsValid()

	assert.NotNil(t, err)
	assert.False(t, isValid)
}

func TestShouldReturnProductId(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	assert.Equal(t, product.GetID(), product.ID)
}

func TestShouldReturnProductName(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	assert.Equal(t, product.GetName(), product.Name)
}

func TestShouldReturnProductPrice(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	assert.Equal(t, product.GetPrice(), product.Price)
}

func TestShouldReturnProductStatus(t *testing.T) {

	product := application.NewProduct("Product 1", 10)

	assert.Equal(t, product.GetStatus(), product.Status)
}
