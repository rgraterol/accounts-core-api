package interfaces

import "github.com/rgraterol/accounts-core-api/domain/entities"

type MovementsService interface {
	P2P(payerID int64, collectorID int64, ammount float64) (entities.Movement, error)
	getPayerAndCollectorAccounts(payerID int64, collectorID int64) (*entities.Account, *entities.Account, error)
}
