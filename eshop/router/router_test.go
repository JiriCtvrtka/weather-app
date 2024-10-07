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
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8888/product/2", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestCreateOrder(t *testing.T) {
	item := models.OrdersCore{
		Username:      "janedoe1",
		Items:         "{\"P432\": 1, \"P433\": 1}",
		Status:        "In Progress",
		Delivery:      "courier",
		DeliveryPrice: 40,
		TotalPrice:    140,
		Currency:      "USD",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8888/order", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestUpdateOrder(t *testing.T) {
	item := models.OrdersCore{
		Items:         "{\"testProd\": 1, \"testProdd\": 3}",
		Status:        "delivered",
		Delivery:      "GLS",
		DeliveryPrice: 55,
		TotalPrice:    251,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8888/order/3", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestDeleterOrder(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8888/order/5", nil)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestCreateUser(t *testing.T) {
	item := models.UserType{
		Firstname:      "user",
		Lastname:       "test",
		Username:       "usertest3",
		Password:       "abc",
		Email:          "usertest@test.com",
		City:           "citytest",
		Street:         "street",
		Number:         "12A",
		AdditionalInfo: "",
		Zipcode:        554422,
		Age:            22,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8888/user", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestDeleterUser(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8888/user/usertest3", nil)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}

func TestUpdateUser(t *testing.T) {
	item := models.UserType{
		Firstname:      "userupdate",
		Lastname:       "test",
		Username:       "usertest3",
		Password:       "abc",
		Email:          "usertest@test.com",
		City:           "citytest",
		Street:         "street",
		Number:         "12A",
		AdditionalInfo: "",
		Zipcode:        554422,
		Age:            22,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(item)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8888/user/usertest3", &buf)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NoError(t, err)

	require.NotEmpty(t, res)
	require.NotEmpty(t, res.Body)
}
