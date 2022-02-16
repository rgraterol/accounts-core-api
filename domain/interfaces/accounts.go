package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type AccountsRepository interface {
	CreateAccount(account entities.Account) (entities.Account, error)
	GetAccountByUserID(userID int64) (entities.Account, error)
	UpdateAccount(account entities.Account) (entities.Account, error)
}

type AccountsService interface {
	SaveNewAccount(id int64, countryID string) error
}
