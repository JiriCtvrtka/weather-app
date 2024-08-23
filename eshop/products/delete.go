package products

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

	product := StaticWay(id)
	if len(product) == 0 {
		return fmt.Errorf("product id %s not found", id)
	}

	_, err = db.Exec("DELETE FROM products WHERE id = $1", id)

	return err
}
