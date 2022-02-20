package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/responses"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
	"github.com/rgraterol/accounts-core-api/domain/services"
)

const (
	defaultUserIDParam = "userID"
)

func MakeMovement(s interfaces.MovementsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input, err := getInputFromRequest(r)
		if err != nil {
			zap.S().Error(err)
			responses.BadRequest(w, err.Error())
			return
		}
		movement, err := s.P2P(*input)
		if err != nil && err == services.NFSError {
			zap.S().Error(err)
			responses.BadRequest(w, err.Error())
			return
		}
		if err != nil {
			zap.S().Error(err)
			responses.Error(w, err)
			return
		}
		responses.OK(w, movement)
		return
	}
}

func MakeDeposit(s interfaces.MovementsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, err := strconv.Atoi(chi.URLParam(r, defaultUserIDParam))
		if err != nil {
			zap.S().Error(err)
			responses.BadRequest(w, "invalid "+defaultUserIDParam)
			return
		}

		deposit, err := getDepositFromRequest(r)
		if err != nil {
			zap.S().Error(err)
			responses.BadRequest(w, err.Error())
			return
		}
		movement, err := s.MakeDeposit(int64(userID), *deposit)
		if err != nil {
			responses.Error(w, err)
			return
		}
		responses.OK(w, movement)
		return
	}
}

func getInputFromRequest(r *http.Request) (*entities.MovementInput, error) {
	var input entities.MovementInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}
	if input.CollectorUserID == 0 {
		return nil, errors.New("collector cannot be null or 0")
	}
	if input.PayerUserID == 0 {
		return nil, errors.New("payer cannot be null or 0")
	}
	return &input, validateDeposit(input.Deposit)
}

func getDepositFromRequest(r *http.Request) (*entities.Deposit, error) {
	var deposit entities.Deposit
	err := json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		return nil, err
	}
	return &deposit, validateDeposit(deposit)
}

func validateDeposit(deposit entities.Deposit) error {
	if deposit.CurrencyID == "" {
		return errors.New("currency cannot be empty")
	}
	if deposit.Amount == float64(0) {
		return errors.New("the amount cannot be zero or null")
	}
	return nil
}
