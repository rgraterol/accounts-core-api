package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
)

var (
	DuplicatedAccountError = errors.New("account already exist")
)

type Accounts struct {
	Repository interfaces.AccountsRepository
}

func (s *Accounts) SaveNewAccount(userID int64, countryID string) error {
	_, err := s.Repository.GetByUserID(userID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = s.Repository.Create(buildNewAccount(userID, countryID))
		return err
	}
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return DuplicatedAccountError
}

func (s *Accounts) Get(userID int64) (entities.Account, error) {
	return s.Repository.GetByUserID(userID)
}

func buildNewAccount(id int64, countryID string) entities.Account {
	return entities.Account{
		UserID:          id,
		Country:         countryID,
		CurrencyID:      entities.CountriesCurrenciesMap[countryID],
		AvailableAmount: float64(0),
	}
}
