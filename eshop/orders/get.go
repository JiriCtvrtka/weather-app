package orders

import (
	"fmt"
	"log"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func StaticWay(id string) []models.OrdersType {
	db, err := db.Connection()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM orders"
	if id != "" {
		query = fmt.Sprintf("%s WHERE id = '%s'", query, id)
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	res := []models.OrdersType{}
	for rows.Next() {
		var order models.OrdersType
		err := rows.Scan(&order.Id, &order.Username, &order.Items, &order.Status, &order.Delivery, &order.DeliveryPrice, &order.TotalPrice, &order.Currency)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, order)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func DynamicWay(table string) []map[string]any {
	fmt.Println(table)
	db, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + table)
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
