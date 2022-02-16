package repositories

import (
	"go.uber.org/zap"

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
