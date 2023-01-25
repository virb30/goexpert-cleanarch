package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderIsValid(t *testing.T) {
	_, err := NewOrder("", 50.0, 10.0)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid id")

	_, err = NewOrder("1", -5.0, 10.0)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid price")

	_, err = NewOrder("1", 5.0, 0)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid tax")

	order, _ := NewOrder("1", 50.0, 10.0)
	err = order.IsValid()
	assert.NoError(t, err)
}

func TestCalculateFinalPrice(t *testing.T) {
	order, err := NewOrder("1", 50.0, 10.0)
	assert.Nil(t, err)
	order.CalculateFinalPrice()
	assert.Equal(t, 60.0, order.FinalPrice)
}
