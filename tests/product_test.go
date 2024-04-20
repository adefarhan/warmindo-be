package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/stretchr/testify/require"
)

var productId string

func TestCreateProducts_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
		"name": "Burjo",
		"price": 5000,
		"stock": 50
	  }`)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusCreated, resp.Code)

	var response struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Data   product.Product `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)

	productId = response.Data.ID
}

func TestGetAllProducts_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/products", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string            `json:"status"`
		Code   int               `json:"code"`
		Data   []product.Product `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestGetDetailProducts_Success(t *testing.T) {
	SetupTest()

	req, err := http.NewRequest("GET", "/products/"+productId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Data   product.Product `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestEditDetailProducts_Success(t *testing.T) {
	SetupTest()

	request := []byte(`{
      "name": "Burjo",
      "price": 4000,
      "stock": 40
	}`)

	req, err := http.NewRequest("PUT", "/products/"+productId, bytes.NewBuffer(request))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Data   product.Product `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}

func TestDeleteProducts_Success(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/products/"+productId, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal("Failed test create product")
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	Router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Data   product.Product `json:"data"`
	}

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	require.Equal(t, "success", response.Status)
}
