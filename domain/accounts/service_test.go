package accounts_test

import (
	"testing"

	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/domain/accounts"
	"github.com/rgraterol/accounts-core-api/infrastructure/init/initializers"
	"github.com/stretchr/testify/assert"
)

const (
	mockCountryID = "UK"
	mockUserID    = int64(1)
)

func init() {
	initializers.MockDatabaseInitializer()
}

func Test_GivenValidInputs_WhenCreate_ThenReturnOk(t *testing.T) {
	// Given
	clearTestDB()
	sut := accounts.Service{}
	// When
	err := sut.SaveNewAccount(mockUserID, mockCountryID)
	// Then
	assert.Nil(t, err)
}

func Test_GivenExistingInputs_WhenCreate_ThenReturnDuplicatedError(t *testing.T) {
	// Given
	clearTestDB()
	sut := accounts.Service{}
	// When
	err := sut.SaveNewAccount(mockUserID, mockCountryID)
	assert.Nil(t, err)
	err = sut.SaveNewAccount(mockUserID, mockCountryID)
	// Then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), accounts.DuplicatedAccountError)
}

func clearTestDB() {
	db.Gorm.Exec("DELETE FROM accounts")
}
