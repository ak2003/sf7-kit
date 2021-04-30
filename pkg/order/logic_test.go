package order

import (
	"context"
	"errors"
	"gt-kit/pkg/order/mocks"
	"gt-kit/pkg/order/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var s service

func TestLogic(t *testing.T) {
	a := s.logicTest(context.Background(), 0)
	assert.Equal(t, a, 0)
}

func TestDeleteItemCart(t *testing.T) {

	param := model.DeleteItemCartRequest{
		CartID:      "1",
		IdxItemCart: 0,
	}
	mk := &mocks.Repository{}
	s := service{repository: mk}
	// mock variable anything, comment 2, comment 3
	mk.On("GetShoppingCart", mock.Anything, mock.Anything).Return(nil, errors.New("test")).Once()

	_, err := s.DeleteItemCart(context.Background(), param)
	assert.Error(t, err)
}
