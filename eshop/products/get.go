package products

import (
	"fmt"
	"log"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func StaticWay(id string) []models.ProductType {
	db, err := db.Connection()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM products"
	if id != "" {
		query = fmt.Sprintf("%s WHERE id = '%s'", query, id)
	}

	rows, err := db.Query(query)
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

func JSONBWay() []map[string]string {
	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products2")
	if err != nil {
		log.Fatal(err)
	}

	res := []map[string]string{}
	for rows.Next() {
		var id string
		var p string
		err := rows.Scan(&id, &p)
		if err != nil {
			log.Fatal(err)
		}

		row := make(map[string]string)
		row[id] = p

		fmt.Println(row[id])
		res = append(res, row)
	}

	return res
}
