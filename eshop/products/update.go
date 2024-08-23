package products

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Update(id string, product models.ProductCore) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	p := StaticWay(id)
	if len(p) == 0 {
		return fmt.Errorf("product id %s not found", id)
	}

	_, err = db.Exec("UPDATE products SET (name, description, currency, count, price) = ($1, $2, $3, $4, $5) WHERE id = $6",
		product.Name, product.Description, product.Currency, product.Count, product.Price, id)

	return err
}
