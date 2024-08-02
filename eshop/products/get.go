package products

import (
	"log"

	"eshop/db"
	"eshop/models"
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

func DynamicWay() []any {
	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	res := []any{}

	for rows.Next() {
		row := []any{}
		for range len(columnTypes) {
			row = append(row, new(string))
		}

		err := rows.Scan(row...)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
