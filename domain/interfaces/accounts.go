package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type AccountsRepository interface {
	Create(account entities.Account) (entities.Account, error)
	GetByUserID(userID int64) (entities.Account, error)
	Update(account entities.Account) (entities.Account, error)
}

type AccountsService interface {
	SaveNewAccount(id int64, countryID string) error
}
