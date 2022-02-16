package services

import (
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
)

type Users struct {
	AccountsService interfaces.AccountsService
}

func (s *Users) ReadUsersFeed(message entities.UserMsg) error {
	if message.Headers.NewUser {
		err := s.AccountsService.SaveNewAccount(message.ID, message.CountryID)
		if err != nil {
			return err
		}
	}
	return nil
}
