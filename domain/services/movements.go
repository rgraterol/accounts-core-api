package services

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
)

var (
	NFSError = errors.New("NFSError")
)

type Movements struct {
	AccountsRepository  interfaces.AccountsRepository
	MovementsRepository interfaces.MovementsRepository
	PayerAccount        *entities.Account
	CollectorAccount    *entities.Account
	Deposit             *entities.Deposit
	Movement            *entities.Movement
}

func (m *Movements) P2P(input entities.MovementInput) (*entities.Movement, error) {
	m.Deposit = &input.Deposit
	err := m.GetPayerAndCollectorAccounts(input)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	if m.PayerAccount.CanMakeOutputMovement(input.Deposit.Amount) {
		// Here we should lock the registries to avoid race conditions to the DB
		// This can be done with a distributed cache or with a DB Lock
		return m.Movement, m.TransferP2P()
	}
	return nil, NFSError
}

func (m *Movements) MakeDeposit(userID int64, deposit entities.Deposit) (*entities.Movement, error) {
	m.Deposit = &deposit
	err := m.GetPayerAndCollectorAccounts(entities.MovementInput{CollectorUserID: userID, PayerUserID: userID})
	if err != nil {
		return nil, err
	}
	m.CollectorAccount.AddAmount(deposit.Amount)
	err = m.buildAndCreateNewMovement()
	if err != nil {
		return nil, err
	}
	return m.Movement, m.SaveDepositWithRollback()
}

func (m *Movements) GetPayerAndCollectorAccounts(input entities.MovementInput) error {
	payerAccount, err := m.AccountsRepository.GetByUserID(input.PayerUserID)
	if err != nil {
		err = errors.Wrap(err, "cannot retrieve payer account")
		return err
	}
	m.PayerAccount = &payerAccount
	collectorAccount, err := m.AccountsRepository.GetByUserID(input.CollectorUserID)
	if err != nil {
		err = errors.Wrap(err, "cannot retrieve collector account")
		return err
	}
	m.CollectorAccount = &collectorAccount
	return nil
}

func (m *Movements) TransferP2P() error {
	m.PayerAccount.DebitAmount(m.Deposit.Amount)
	m.CollectorAccount.AddAmount(m.Deposit.Amount)
	err := m.buildAndCreateNewMovement()
	if err != nil {
		return err
	}
	return m.SaveTransferWithRollback()
}

func (m *Movements) SaveTransferWithRollback() error {
	err := m.SavePayerAccountWithRollback(nil)
	if err != nil {
		return err
	}
	err = m.SaveCollectorAccountWithRollback()
	if err != nil {
		return err
	}
	err = m.UpdateMovement(entities.MovementDone, nil)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movements) SaveDepositWithRollback() error {
	err := m.SaveCollectorAccountWithRollback()
	if err != nil {
		return err
	}
	err = m.UpdateMovement(entities.MovementDone, nil)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movements) UpdateMovement(status string, falledErr error) error {
	m.Movement.Status = status
	_, err := m.MovementsRepository.Update(*m.Movement)
	if err != nil {
		err = errors.Wrap(err, "cannot update movement "+status)
		zap.S().Error(err)
		return err
	}
	if falledErr != nil {
		err = errors.Wrap(falledErr, "cannot complete movement "+status)
		zap.S().Error(falledErr)
	}
	return falledErr
}

func (m *Movements) SavePayerAccountWithRollback(falledError error) error {
	_, err := m.AccountsRepository.Update(*m.PayerAccount)
	if err != nil {
		// ROLLING BACK MOVEMENT TO ERROR STATUS
		zap.S().Error(err)
		return m.UpdateMovement(entities.MovementErrorSavingAccount, err)
	}
	return falledError
}

func (m *Movements) SaveCollectorAccountWithRollback() error {
	_, err := m.AccountsRepository.Update(*m.CollectorAccount)
	if err != nil {
		// ROLLING BACK amount FOR PAYER
		zap.S().Error(err)
		m.PayerAccount.AddAmount(m.Deposit.Amount)
		return m.SavePayerAccountWithRollback(err)
	}
	return nil
}

func (m *Movements) BuildNewMovement() entities.Movement {
	return entities.Movement{
		PayerUserID:        m.PayerAccount.UserID,
		PayerAccountID:     m.PayerAccount.ID,
		CollectorUserID:    m.CollectorAccount.UserID,
		CollectorAccountID: m.CollectorAccount.ID,
		Amount:             m.Deposit.Amount,
		PayerBalance:       m.PayerAccount.AvailableAmount,
		CollectorBalance:   m.CollectorAccount.AvailableAmount,
		Reason:             m.Deposit.Reason,
		CurrencyID:         m.Deposit.CurrencyID,
		CountryID:          m.PayerAccount.Country,
		Status:             entities.MovementInProgress,
	}
}

func (m *Movements) buildAndCreateNewMovement() error {
	movement := m.BuildNewMovement()
	createdMovement, err := m.MovementsRepository.Create(movement)
	if err != nil {
		err = errors.Wrap(err, "cannot save movement")
		zap.S().Error(err)
		return err
	}
	m.Movement = &createdMovement
	return nil
}
