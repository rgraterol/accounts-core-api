package users

import (
	"github.com/rgraterol/accounts-core-api/domain/accounts"
)

type Service struct {
	AccountService accounts.Interface
}

func (s *Service) ReadUsersFeed(message UserMsg) error {
	if message.Headers.NewUser {
		err := s.AccountService.SaveNewAccount(message.ID, message.CountryID)
		if err != nil {
			return err
		}
	}
	return nil
}
