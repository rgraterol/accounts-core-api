package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type MovementsService interface {
	P2P(input entities.MovementInput) (*entities.Movement, error)
	GetPayerAndCollectorAccounts(input entities.MovementInput) error
	TransferP2P() error
	SaveTransferWithRollback() error
	UpdateMovement(status string) error
	SavePayerAccountWithRollback() error
	SaveCollectorAccountWithRollback() error
	BuildNewMovement() entities.Movement
}

type MovementsRepository interface {
	Create(movement entities.Movement) (entities.Movement, error)
	Update(movement entities.Movement) (entities.Movement, error)
}
