package users_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/users"
	"github.com/stretchr/testify/assert"
)

const (
	accountServiceError = "error on accounts service"
)

func Test_GivenNewUser_WhenErrorOnAccountsService_ThenReturnError(t *testing.T) {
	// GIVEN
	msg := buildMsgNewUser()
	sut := users.Service{
		AccountService: &AccountServiceMockError{},
	}

	// WHEN
	err := sut.ReadUsersFeed(msg)

	// THEN
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), accountServiceError)
}

func Test_GivenMsgWithoutNewUser_ThenReturnOk(t *testing.T) {
	// GIVEN
	msg := buildMsgExistingUser()
	sut := users.Service{
		AccountService: &AccountServiceMockOk{},
	}

	// WHEN
	err := sut.ReadUsersFeed(msg)

	// THEN
	assert.Nil(t, err)
}

func Test_GivenNewUser_WhenAccountsOk_ThenReturnOk(t *testing.T) {
	// GIVEN
	msg := buildMsgExistingUser()
	sut := users.Service{}

	// WHEN
	err := sut.ReadUsersFeed(msg)

	// THEN
	assert.Nil(t, err)
}

func buildMsgNewUser() users.UserMsg {
	return users.UserMsg{
		ID: 1,
		Headers: users.UserMsgHeaders{
			NewUser: true,
		},
	}
}

func buildMsgExistingUser() users.UserMsg {
	return users.UserMsg{
		ID: 1,
	}
}

type AccountServiceMockOk struct{}

func (s *AccountServiceMockOk) SaveNewAccount(id int64, countryID string) error {
	return nil
}

type AccountServiceMockError struct{}

func (s *AccountServiceMockError) SaveNewAccount(id int64, countryID string) error {
	return errors.New(accountServiceError)
}
