package repositories_test

import (
	"testing"

	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/application/repositories"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/infrastructure/initializers"
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
	clearAccountsTableDB()
	sut := repositories.Accounts{}
	// When
	account, err := sut.Create(buildMockAccount())
	// Then
	assert.Nil(t, err)
	assert.Equal(t, mockCurrency, account.CurrencyID)
}

func Test_GivenValidId_WhenGetByUserID_ThenReturnAccount(t *testing.T) {
	// Given
	clearAccountsTableDB()
	sut := repositories.Accounts{}
	mockAccount := buildMockAccount()
	_, err := sut.Create(mockAccount)
	assert.Nil(t, err)
	// When
	account, err := sut.GetByUserID(mockAccount.UserID)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, mockCurrency, account.CurrencyID)
}

func Test_GivenValidAccount_WhenUpdate_ThenReturnAccount(t *testing.T) {
	// Given
	clearAccountsTableDB()
	sut := repositories.Accounts{}
	mockAccount := buildMockAccount()
	_, err := sut.Create(mockAccount)
	assert.Nil(t, err)
	// When
	mockAccount.AvailableAmount = 100000
	_, err = sut.Update(mockAccount)
	// Then
	assert.Nil(t, err)
	account, err := sut.GetByUserID(mockAccount.UserID)
	assert.Equal(t, float64(100000), account.AvailableAmount)
}

func buildMockAccount() entities.Account {
	return entities.Account{
		ID:         1,
		UserID:     mockUserID,
		CurrencyID: mockCurrency,
	}
}

func clearAccountsTableDB() {
	db.Gorm.Exec("DELETE FROM accounts")
}
