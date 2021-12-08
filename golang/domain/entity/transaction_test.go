package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900

	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1200

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "You dont have limit for this transaction", err.Error())
}

func TestTransaction_IsValidWithAmountLessThen1(t *testing.T) {
	transaction := NewTransaction()

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "The amount must to be greater than 1", err.Error())
}
