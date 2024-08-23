package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weather-app/eshop/models"
)

func TestCreateProduct(t *testing.T) {
	item := models.ProductCore{
		Name:        "name",
		Description: "desc",
		Currency:    "EUR",
		Count:       5,
		Price:       40,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8888/product", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestDeleteProduct(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8888/product/1", nil)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestUpdateProduct(t *testing.T) {
	item := models.ProductCore{
		Name:        "xxx",
		Description: "yyyy",
		Currency:    "USD",
		Count:       51,
		Price:       75,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8888/product/1", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}
