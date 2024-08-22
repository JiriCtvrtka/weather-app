package products

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Create(item models.CreateProduct) (int64, error) {
	db, err := db.Connection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer db.Close()

	var id int64
	err = db.QueryRow(
		"INSERT INTO products (name, description, currency, count, price) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		item.Name, item.Description, item.Currency, item.Count, item.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
