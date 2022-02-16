package services_test

import (
	"testing"

	"github.com/rgraterol/accounts-core-api/domain/services"
	"github.com/stretchr/testify/assert"
)

const (
	mockCountryID   = "UK"
	mockUserID      = int64(1)
	mockCollectorID = int64(2)
)

func Test_GivenValidInputs_WhenCreate_ThenReturnOk(t *testing.T) {
	// Given
	sut := services.Accounts{}
	// When
	err := sut.SaveNewAccount(mockUserID, mockCountryID)
	// Then
	assert.Nil(t, err)
}

func Test_GivenExistingInputs_WhenCreate_ThenReturnDuplicatedError(t *testing.T) {
	// Given
	sut := services.Accounts{}
	// When
	err := sut.SaveNewAccount(mockUserID, mockCountryID)
	assert.Nil(t, err)
	err = sut.SaveNewAccount(mockUserID, mockCountryID)
	// Then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), services.DuplicatedAccountError)
}
