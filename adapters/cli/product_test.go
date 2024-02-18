package cli_test

import (
	"testing"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	productName := "Product Test"
	producPrice := 10.0
	productStatus := "enabled"
	productId := "abc"

	pMock := mock_application.NewMockProductInterface(ctrl)
	pMock.EXPECT().GetID().Return(productId).AnyTimes()
	pMock.EXPECT().GetPrice().Return(producPrice).AnyTimes()
	pMock.EXPECT().GetName().Return(productName).AnyTimes()
	pMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	s := mock_application.NewMockProductServiceInterface(ctrl)

	s.EXPECT().Create(productName, producPrice).Return(pMock, nil).AnyTimes()
	s.EXPECT().Get(productId).Return(pMock, nil).AnyTimes()
	s.EXPECT().Enable(gomock.Any()).Return(pMock, nil).AnyTimes()
	s.EXPECT().Enable(gomock.Any()).Return(pMock, nil).AnyTimes()

	r, err := cli.Run(s, "create", "", productName, producPrice)

	assert.Nil(t, err)
	assert.True(t, len(r) > 0)

	r, err = cli.Run(s, "disabled", productId, "", 0)

	assert.Nil(t, err)
	assert.True(t, len(r) > 0)

	r, err = cli.Run(s, "enable", productId, "", 0)

	assert.Nil(t, err)
	assert.True(t, len(r) > 0)

	r, err = cli.Run(s, "get", productId, "", 0)

	assert.Nil(t, err)
	assert.True(t, len(r) > 0)
}
