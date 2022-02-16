package handlers_test

import (
	"github.com/rgraterol/accounts-core-api/domain/entities"
)

type MovementsMock struct {
}

func (m *MovementsMock) P2P(input entities.MovementInput) (*entities.Movement, error) {
	return nil, nil
}

func (m *MovementsMock) GetPayerAndCollectorAccounts(input entities.MovementInput) error {
	return nil
}

func (m *MovementsMock) TransferP2P() error {
	return nil
}

func (m *MovementsMock) SaveTransferWithRollback() error {
	return nil
}

func (m *MovementsMock) UpdateMovement(status string) error {
	return nil
}

func (m *MovementsMock) SavePayerAccountWithRollback() error {
	return nil
}

func (m *MovementsMock) SaveCollectorAccountWithRollback() error {
	return nil
}

func (m *MovementsMock) buildNewMovement() entities.Movement {
	return entities.Movement{}
}
