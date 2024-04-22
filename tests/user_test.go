package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/stretchr/testify/require"
)

var customerId string

func TestCreateCustomers_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
		"name": "Ade Farhan",
		"phoneNumber": "9218412423",
		"address": "jakarta"
	  }`)

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test create customer")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusCreated, resp.Code)

	var response struct {
		Status string            `json:"status"`
		Code   int               `json:"code"`
		Data   customer.Customer `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)

	customerId = response.Data.ID
}

func TestGetAllCustomers_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/customers", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test get all customer")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string              `json:"status"`
		Code   int                 `json:"code"`
		Data   []customer.Customer `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestGetDetailCustomers_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/customers/"+customerId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test get detail customer")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string            `json:"status"`
		Code   int               `json:"code"`
		Data   customer.Customer `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestEditCustomers_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
		"name": "Ade Farhan Edit",
		"phoneNumber": "1111111111",
		"address": "sumatera"
	  }`)
	req, err := http.NewRequest("PUT", "/customers/"+customerId, bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test update customer")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string            `json:"status"`
		Code   int               `json:"code"`
		Data   customer.Customer `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestDeleteCustomers_Success(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/"+customerId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test delet customer")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string            `json:"status"`
		Code   int               `json:"code"`
		Data   customer.Customer `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}
