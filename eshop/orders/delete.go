package orders

import (
	"fmt"

	"github.com/weather-app/eshop/db"
)

func Delete(id string) error {
	db, err := db.Connection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	order := StaticWay(id)
	if len(order) == 0 {
		return fmt.Errorf("order id %s not found", id)
	}

	_, err = db.Exec("DELETE FROM orders WHERE id = $1", id)

	return err
}
