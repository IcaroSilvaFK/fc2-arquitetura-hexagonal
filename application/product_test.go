package application_test

import (
	"testing"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
)

func TestShouldEnableProductWhenPriceIsGTOnZero(t *testing.T) {

	product := application.NewProduct("Product 1", "enabled", 10)

	if err := product.Enable(); err != nil {
		t.Error("Expected the product to be enable, but got error")
	}

}

func TestShouldNotEnableProductWhenPriceIsLEOnZero(t *testing.T) {

	product := application.NewProduct("Product 1", "enabled", 0)

	if err := product.Enable(); err == nil {
		t.Error("The product should not be enable")
	}

}
