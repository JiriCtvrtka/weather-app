package utils

import (
	"fmt"
	"log"

	"github.com/weather-app/eshop/db"
)

func DynamicWay(table string) []map[string]any {
	fmt.Println("dinamic")
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
