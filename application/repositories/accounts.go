package repositories

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/domain/entities"
)

type Accounts struct{}

func (s *Accounts) GetAccountByUserID(userID int64) (entities.Account, error) {
	var account entities.Account
	trx := db.Gorm.Where("user_id = ?", userID).First(&account)
	return account, trx.Error
}

func (s *Accounts) CreateAccount(account entities.Account) (entities.Account, error) {
	trx := db.Gorm.Create(&account)
	if trx.Error != nil {
		zap.S().Error(trx.Error)
		return entities.Account{}, trx.Error
	}
	return account, nil
}

// Here we should only update single column, I'm updating whole object because is a test app
func (s *Accounts) UpdateAccount(account entities.Account) (entities.Account, error) {
	trx := db.Gorm.Save(&account)
	if trx.Error != nil {
		err := errors.Wrap(trx.Error, "cannot update account")
		zap.S().Error(err)
		return account, err
	}
	return account, nil
}
