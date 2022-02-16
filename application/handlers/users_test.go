package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/handlers"
	"github.com/rgraterol/accounts-core-api/domain/users"
	"github.com/stretchr/testify/assert"
)

const (
	applicationJsonHeader = "application/json"
	messageKey            = "message"
	serviceErrorMessage   = "cannot save new user"
)

func Test_GivenUsersNews_WhenEmptyBody_ThenReturnBadRequest(t *testing.T) {
	//GIVEN
	ts := httptest.NewServer(http.HandlerFunc(handlers.UsersNews(&ServiceMockOk{})))
	defer ts.Close()
	//WHEN
	res, _ := http.Post(ts.URL, applicationJsonHeader, nil)
	var resp map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&resp)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, "EOF", resp[messageKey])
}

func Test_GivenUsersNews_WhenNilUserID_ThenReturnBadRequest(t *testing.T) {
	//GIVEN
	ts := httptest.NewServer(http.HandlerFunc(handlers.UsersNews(&ServiceMockOk{})))
	defer ts.Close()
	body, err := json.Marshal(buildMessageEmptyID())
	assert.Nil(t, err)
	//WHEN
	res, _ := http.Post(ts.URL, applicationJsonHeader, bytes.NewBuffer(body))
	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, resp[messageKey], "userID cannot be empty")
}

func Test_GivenUsersNews_WhenInvalidUserID_ThenReturnBadRequest(t *testing.T) {
	//GIVEN
	ts := httptest.NewServer(http.HandlerFunc(handlers.UsersNews(&ServiceMockOk{})))
	defer ts.Close()

	body, err := json.Marshal(buildMessageInvalidID())
	assert.Nil(t, err)
	//WHEN
	res, _ := http.Post(ts.URL, applicationJsonHeader, bytes.NewBuffer(body))
	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, resp[messageKey], "cannot unmarshal string into Go struct")
}

func Test_GivenUsersNews_WhenErrorOnService_ThenReturnInternalServerError(t *testing.T) {
	//GIVEN
	ts := httptest.NewServer(http.HandlerFunc(handlers.UsersNews(&ServiceMockError{})))
	defer ts.Close()
	body, err := json.Marshal(buildValidMessage())
	assert.Nil(t, err)
	//WHEN
	res, _ := http.Post(ts.URL, applicationJsonHeader, bytes.NewBuffer(body))
	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Contains(t, serviceErrorMessage, resp[messageKey])
}

func Test_GivenUsersNews_WhenValidUserID_ThenOk(t *testing.T) {
	//GIVEN
	ts := httptest.NewServer(http.HandlerFunc(handlers.UsersNews(&ServiceMockOk{})))
	defer ts.Close()
	body, err := json.Marshal(buildValidMessage())
	assert.Nil(t, err)
	//WHEN
	res, _ := http.Post(ts.URL, applicationJsonHeader, bytes.NewBuffer(body))
	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	//THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Contains(t, "users news processed ok", resp[messageKey])
}

type ServiceMockOk struct{}

func (s *ServiceMockOk) ReadUsersFeed(message users.UserMsg) error {
	return nil
}

type ServiceMockError struct{}

func (s *ServiceMockError) ReadUsersFeed(message users.UserMsg) error {
	return errors.New(serviceErrorMessage)
}

func buildMessageEmptyID() map[string]interface{} {
	return map[string]interface{}{
		"msg": map[string]interface{}{
			"id": 0,
		},
	}
}

func buildMessageInvalidID() map[string]interface{} {
	return map[string]interface{}{
		"msg": map[string]interface{}{
			"id": "not_int",
		},
	}
}

func buildValidMessage() map[string]interface{} {
	return map[string]interface{}{
		"msg": map[string]interface{}{
			"id": 100,
		},
	}
}
