package application_test

import (
	"testing"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	mock_application "github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetProductById(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	r, err := service.Get("abc")

	assert.Nil(t, err)
	assert.Equal(t, product, r)

}

func TestShouldCreateNewProduct(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	svc := application.ProductService{Persistence: persistence}

	r, err := svc.Create("Product 1", 10)

	assert.Nil(t, err)
	assert.Equal(t, r, product)
}

func TestShouldEnableProduct(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	svc := application.ProductService{Persistence: persistence}

	r, err := svc.Enable(product)

	assert.Nil(t, err)
	assert.Equal(t, r, product)

	r, err = svc.Disable(product)

	assert.Nil(t, err)
	assert.Equal(t, r, product)
}
