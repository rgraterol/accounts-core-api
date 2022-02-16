package repositories_test

import (
	"testing"

	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/application/repositories"
	"github.com/rgraterol/accounts-core-api/domain/accounts"
	"github.com/rgraterol/accounts-core-api/infrastructure/init/initializers"
	"github.com/stretchr/testify/assert"
)

const (
	mockCurrency = "EUR"
	mockUserID   = int64(1)
)

func init() {
	initializers.MockDatabaseInitializer()
}

func Test_GivenValidAccount_WhenCreateAccount_ThenReturnOk(t *testing.T) {
	// Given
	clearTestDB()
	sut := repositories.Accounts{}
	// When
	account, err := sut.CreateAccount(buildMockAccount())
	// Then
	assert.Nil(t, err)
	assert.Equal(t, mockCurrency, account.CurrencyID)
}

func buildMockAccount() accounts.Account {
	return accounts.Account{
		ID:         1,
		UserID:     mockUserID,
		CurrencyID: mockCurrency,
	}
}

func clearTestDB() {
	db.Gorm.Exec("DELETE FROM accounts")
}
