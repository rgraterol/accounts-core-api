package services_test

import (
	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/entities"
)

const (
	genericDBError = "error connecting to DB"
)

var errorDB = errors.New(genericDBError)

type AccountsRepositoryMockOk struct{}

func (r *AccountsRepositoryMockOk) CreateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, nil
}

func (r *AccountsRepositoryMockOk) GetAccountByUserID(userID int64) (entities.Account, error) {
	return entities.Account{}, nil
}
func (r *AccountsRepositoryMockOk) UpdateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, nil
}

type AccountsRepositoryMockPayerRich struct{}

func (r *AccountsRepositoryMockPayerRich) CreateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, nil
}

func (r *AccountsRepositoryMockPayerRich) GetAccountByUserID(userID int64) (entities.Account, error) {
	return entities.Account{AvailableAmount: 10000}, nil
}
func (r *AccountsRepositoryMockPayerRich) UpdateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, nil
}

type AccountsRepositoryMockError struct{}

func (r *AccountsRepositoryMockError) CreateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, errorDB
}

func (r *AccountsRepositoryMockError) GetAccountByUserID(userID int64) (entities.Account, error) {
	return entities.Account{}, errorDB
}
func (r *AccountsRepositoryMockError) UpdateAccount(account entities.Account) (entities.Account, error) {
	return entities.Account{}, errorDB
}

type MovementsRepositoryMockOk struct{}

func (m *MovementsRepositoryMockOk) Create(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, nil
}
func (m *MovementsRepositoryMockOk) Update(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, nil
}

type MovementsRepositoryMockError struct{}

func (m *MovementsRepositoryMockError) Create(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, errorDB
}
func (m *MovementsRepositoryMockError) Update(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, errorDB
}
