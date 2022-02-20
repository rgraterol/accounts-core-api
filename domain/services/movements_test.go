package services_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/services"
	"github.com/stretchr/testify/assert"
)

const (
	amountMock = float64(100)
)

func TestBuildNewMovement(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	//WHEN
	movement := sut.BuildNewMovement()
	//THEN
	assert.Equal(t, amountMock, movement.Amount)
}

func Test_GivenRepositoryOk_WhenSaveCollector_ThenOK(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	//WHEN
	err := sut.SaveCollectorAccountWithRollback()
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, mockCollectorID, sut.CollectorAccount.ID)
}

func Test_GivenRepositoryError_WhenSaveCollector_ThenReturnError(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.AccountsRepository = &AccountsRepositoryMockError{}
	//WHEN
	err := sut.SaveCollectorAccountWithRollback()
	//THEN
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), genericDBError)
}

func Test_GivenRepositoryOk_WhenSavePayer_ThenOK(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	//WHEN
	err := sut.SavePayerAccountWithRollback(nil)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, mockUserID, sut.PayerAccount.ID)
}

func Test_GivenRepositoryOk_WhenSavePayerWithFallenError_ThenReturnError(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	falledErr := errors.New("could not save collector, rollbacked payer")
	//WHEN
	err := sut.SavePayerAccountWithRollback(falledErr)
	//THEN
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "could not save collector, rollbacked payer")
}

func Test_GivenRepositoryError_WhenSavePayer_ThenReturnError(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.AccountsRepository = &AccountsRepositoryMockError{}
	//WHEN
	err := sut.SavePayerAccountWithRollback(nil)
	//THEN
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), genericDBError)
}

func Test_GivenRepositoryError_WhenUpdateMovement_ThenReturnError(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.MovementsRepository = &MovementsRepositoryMockError{}
	//WHEN
	err := sut.UpdateMovement(entities.MovementInProgress, nil)
	//THEN
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), genericDBError)
}

func Test_GivenRepositoryOk_WhenUpdateMovementOk_ThenReturnOk(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.MovementsRepository = &MovementsRepositoryMockOk{}
	//WHEN
	err := sut.UpdateMovement(entities.MovementInProgress, nil)
	//THEN
	assert.Nil(t, err)
}

func Test_GivenAccountRepositoryErr_WhenSaveTransferWithRollback_ThenReturnErr(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.AccountsRepository = &AccountsRepositoryMockError{}
	//WHEN
	err := sut.SaveTransferWithRollback()
	//THEN
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), genericDBError)
}

func Test_GivenMovementRepositoryErr_WhenSaveTransferWithRollback_ThenReturnErr(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.MovementsRepository = &MovementsRepositoryMockError{}
	//WHEN
	err := sut.SaveTransferWithRollback()
	//THEN
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot update movement")
}

func Test_GivenMovementRepositoryErr_WhenSaveTransferP2P_ThenReturnErr(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.MovementsRepository = &MovementsRepositoryMockError{}
	//WHEN
	err := sut.TransferP2P()
	//THEN
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot save movement")
}

func Test_GivenRepositoriesOk_WhenSaveTransferP2P_ThenReturnOk(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	//WHEN
	err := sut.TransferP2P()
	//THEN
	assert.Nil(t, err)
}

func Test_GivenAccountRepositoryErr_WhenGetPayerAndCollectorAccounts_ThenReturnErr(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	sut.AccountsRepository = &AccountsRepositoryMockError{}
	input := buildMockInput()
	//WHEN
	err := sut.GetPayerAndCollectorAccounts(*input)
	//THEN
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot retrieve payer account")
}

func Test_GivenAccountRepositoryOk_WhenGetPayerAndCollectorAccounts_ThenReturnOk(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	input := buildMockInput()
	//WHEN
	err := sut.GetPayerAndCollectorAccounts(*input)
	//THEN
	assert.Nil(t, err)
	// Expect a clean payeraccount because the mock returns an empty struck
	assert.Equal(t, int64(0), sut.PayerAccount.ID)
}

func Test_GivenAccountRepositoryErr_WhenP2P_ThenReturnErr(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	input := buildMockInput()
	sut.AccountsRepository = &AccountsRepositoryMockError{}
	//WHEN
	movement, err := sut.P2P(*input)
	//THEN
	assert.NotNil(t, err)
	assert.Nil(t, movement)
}

func Test_GivenAccountWithoutFunds_WhenP2P_ThenReturnNSFError(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	input := buildMockInput()
	sut.PayerAccount.AvailableAmount = float64(0)
	//WHEN
	movement, err := sut.P2P(*input)
	//THEN
	assert.NotNil(t, err)
	assert.Nil(t, movement)
	assert.Equal(t, services.NFSError, err)
}

func Test_GivenAccountOk_WhenP2P_ThenReturnOk(t *testing.T) {
	//GIVEN
	sut := buildMovementServiceMockOk()
	input := buildMockInput()
	sut.AccountsRepository = &AccountsRepositoryMockPayerRich{}
	//WHEN
	movement, err := sut.P2P(*input)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, amountMock, movement.Amount)
	assert.Equal(t, float64(9900), movement.PayerBalance)
	assert.Equal(t, float64(10100), movement.CollectorBalance)
	assert.Equal(t, "lunch", movement.Reason)
}

func buildMovementServiceMockOk() services.Movements {
	return services.Movements{
		AccountsRepository:  &AccountsRepositoryMockOk{},
		MovementsRepository: &MovementsRepositoryMockOk{},
		Deposit:             buildMockInput(),
		PayerAccount:        buildMockPayerAccount(),
		CollectorAccount:    buildMockCollectorAccount(),
		Movement:            buildMockMovement(),
	}
}

func buildMockInput() *entities.MovementInput {
	deposit := entities.Deposit{
		Amount:     amountMock,
		Reason:     "lunch",
		CurrencyID: "EUR",
	}
	return &entities.MovementInput{
		Deposit:         deposit,
		PayerUserID:     mockUserID,
		CollectorUserID: mockCollectorID,
	}
}

func buildMockPayerAccount() *entities.Account {
	return &entities.Account{
		ID:              mockUserID,
		UserID:          mockUserID,
		Country:         mockCountryID,
		AvailableAmount: 300,
	}
}

func buildMockCollectorAccount() *entities.Account {
	return &entities.Account{
		ID:              mockCollectorID,
		UserID:          mockCollectorID,
		Country:         mockCountryID,
		AvailableAmount: 0,
	}
}

func buildMockMovement() *entities.Movement {
	return &entities.Movement{
		Status: entities.MovementInProgress,
	}
}
