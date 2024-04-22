package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adefarhan/warmindo-be/internal/entity/user"
	"github.com/stretchr/testify/require"
)

var userId string

func TestCreateUsers_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
		"name": "Ade Farhan",
		"phoneNumber": "9218412423",
		"address": "jakarta"
	  }`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test create user")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusCreated, resp.Code)

	var response struct {
		Status string    `json:"status"`
		Code   int       `json:"code"`
		Data   user.User `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)

	userId = response.Data.ID
}

func TestGetAllUsers_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/users", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create user")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string      `json:"status"`
		Code   int         `json:"code"`
		Data   []user.User `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestGetDetailUsers_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/users/"+userId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create user")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string    `json:"status"`
		Code   int       `json:"code"`
		Data   user.User `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestEditUsers_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
		"name": "Ade Farhan Edit",
		"phoneNumber": "1111111111",
		"address": "sumatera"
	  }`)
	req, err := http.NewRequest("PUT", "/users/"+userId, bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test update user")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string    `json:"status"`
		Code   int       `json:"code"`
		Data   user.User `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestDeleteUsers_Success(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/"+userId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string    `json:"status"`
		Code   int       `json:"code"`
		Data   user.User `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}
