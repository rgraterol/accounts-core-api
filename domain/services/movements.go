package services

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
)

type Service struct {
	AccountsRepository interfaces.AccountsRepository
}

func (s *Service) P2P(payerID int64, collectorID int64, ammount float64) (*entities.Movement, error) {
	payerAccount, _, err := s.getPayerAndCollectorAccounts(payerID, collectorID)
	if err != nil {
		zap.S().Error()
		return nil, err
	}
	if payerAccount.CanMakeOutputMovement(ammount) {
		return nil, nil
	}
	return nil, nil
}

func (s *Service) getPayerAndCollectorAccounts(payerID int64, collectorID int64) (*entities.Account, *entities.Account, error) {
	payerAccount, err := s.AccountsRepository.GetAccountByUserID(payerID)
	if err != nil {
		err = errors.Wrap(err, "cannot retrieve payer account")
		return nil, nil, err
	}

	collectorAccount, err := s.AccountsRepository.GetAccountByUserID(collectorID)
	if err != nil {
		err = errors.Wrap(err, "cannot retrieve collector account")
		return nil, nil, err
	}
	return &payerAccount, &collectorAccount, nil
}
