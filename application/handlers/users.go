package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/responses"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/domain/interfaces"
	"github.com/rgraterol/accounts-core-api/domain/services"
)

// UsersNews this endpoint is a consumer of users-api feed, to create a new account of users recently registered
func UsersNews(s interfaces.UsersService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		message, err := decodeUsersNewsMessage(r)
		if err != nil {
			zap.S().Error(err)
			responses.BadRequest(w, err.Error())
			return
		}
		err = s.ReadUsersFeed(message.Msg)
		// The response we respond here depends on the queue used
		// If we use Google Pub/Sub or ApachePulsar we should not return a 400 status code
		// That's why we're returning an OK message, but we are logging inside
		if err != nil && err == services.DuplicatedAccountError {
			responses.OK(w, err.Error())
			return
		}

		if err != nil {
			zap.S().Error(err)
			responses.Error(w, err)
			return
		}
		responses.OK(w, map[string]string{
			"message": "users news processed ok",
		})
	}
}

func decodeUsersNewsMessage(r *http.Request) (*entities.UsersFeedMessage, error) {
	var message entities.UsersFeedMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		return nil, err
	}
	if message.Msg.ID == 0 {
		return nil, errors.New("userID cannot be empty")
	}
	return &message, nil
}
