package orders

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Update(id string, order models.OrdersCore) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	p := StaticWay(id)
	if len(p) == 0 {
		return fmt.Errorf("order id %s not found", id)
	}

	_, err = db.Exec("UPDATE orders SET (items, status, delivery, delivery_price, total_price) = ($1, $2, $3, $4, $5) WHERE id = $6",
		order.Items, order.Status, order.Delivery, order.DeliveryPrice, order.TotalPrice, id)

	return err
}
