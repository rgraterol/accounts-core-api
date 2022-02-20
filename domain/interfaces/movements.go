package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type MovementsService interface {
	P2P(input entities.MovementInput) (*entities.Movement, error)
	MakeDeposit(userID int64, deposit entities.Deposit) (*entities.Movement, error)
	GetPayerAndCollectorAccounts(input entities.MovementInput) error
	TransferP2P() error
	SaveTransferWithRollback() error
	UpdateMovement(status string, falledErr error) error
	SavePayerAccountWithRollback(falledError error) error
	SaveCollectorAccountWithRollback() error
	BuildNewMovement() entities.Movement
}

type MovementsRepository interface {
	Create(movement entities.Movement) (entities.Movement, error)
	Update(movement entities.Movement) (entities.Movement, error)
}
