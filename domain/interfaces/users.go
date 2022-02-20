package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type UsersService interface {
	ReadUsersFeed(message entities.UserMsg) error
	Get(userID int64) (entities.Account, error)
}
