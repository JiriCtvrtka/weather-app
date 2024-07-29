package router

import (
	"encoding/json"
	"eshop/db"
	"eshop/models"
	"log"
	"net/http"
)

func getAllProducts() []models.ProductType {
	db, err := db.Connection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal(err)
	}

	products := []models.ProductType{}

	for rows.Next() {
		var prod models.ProductType
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Description, &prod.Currency, &prod.Count, &prod.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prod)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return products

}

func Routing() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := getAllProducts()
		bytes, err := json.Marshal(products)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})
}
