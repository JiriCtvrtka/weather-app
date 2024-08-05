package products

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func StaticWay() []models.ProductType {
	db, err := db.Connection()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	res := []models.ProductType{}
	for rows.Next() {
		var prod models.ProductType
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Description, &prod.Currency, &prod.Count, &prod.Price)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, prod)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func DynamicWay() []map[string]any {
	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	res := []map[string]any{}
	for rows.Next() {
		row := []any{}
		for range len(columns) {
			row = append(row, new(string))
		}

		err := rows.Scan(row...)
		if err != nil {
			log.Fatal(err)
		}

		resRow := map[string]any{}
		for i, columnName := range columns {
			resRow[columnName] = row[i]
		}

		res = append(res, resRow)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

type product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
	Status      int     `json:"status,omitempty"`
}

func (p *product) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)

}

func JSONBWay() []map[string]product {
	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products2")
	if err != nil {
		log.Fatal(err)
	}

	res := []map[string]product{}
	for rows.Next() {
		var id string
		var p product
		err := rows.Scan(&id, &p)
		if err != nil {
			log.Fatal(err)
		}

		row := make(map[string]product)
		row[id] = p

		res = append(res, row)
	}

	return res
}
