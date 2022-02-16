package accounts

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/domain/countries"
)

var (
	DuplicatedAccountError = errors.New("account already exist")
)

type Service struct{}

func (s *Service) SaveNewAccount(id int64, countryID string) error {
	var account Account
	trx := db.Gorm.First(&account, id)
	if trx.Error != nil && errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		return createNewAccount(id, countryID)
	}
	if trx.Error != nil {
		zap.S().Error(trx.Error)
		return trx.Error
	}
	return DuplicatedAccountError
}
func createNewAccount(id int64, countryID string) error {
	newAccount := buildNewAccount(id, countryID)
	trx := db.Gorm.Create(&newAccount)
	if trx.Error != nil {
		zap.S().Error(trx.Error)
		return trx.Error
	}
	return nil
}

func buildNewAccount(id int64, countryID string) Account {
	return Account{
		UserID:            id,
		Country:           countryID,
		CurrencyID:        countries.CurrenciesMap[countryID],
		TotalAmount:       float64(0),
		UnavailableAmount: float64(0),
	}
}
