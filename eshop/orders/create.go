package orders

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Create(item models.OrdersCore) (int64, error) {
	db, err := db.Connection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer db.Close()

	var id int64
	err = db.QueryRow(
		"INSERT INTO orders (username, items, status, delivery, delivery_price, total_price, currency) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		item.Username, item.Items, item.Status, item.Delivery, item.DeliveryPrice, item.TotalPrice, item.Currency).Scan(&id)
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
