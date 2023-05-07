package entity_test

import (
	"testing"

	"github.com/fabiofa8/pfa-go/internal/order/entity"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order:= entity.Order{}
	
	assert.Error(t, order.IsValid(), "Invalid Id")
}
func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order:= entity.Order{ID: "123"}
	
	assert.Error(t, order.IsValid(), "invalid price")
}
func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order:= entity.Order{ID: "123", Price: 10}
	
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	
	assert.NoError(t, err)
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 2.0)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()

	assert.NoError(t, err)
	assert.Equal(t, order.FinalPrice, 12.0)
}