package router

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/weather-app/eshop/models"
	"github.com/weather-app/eshop/products"
)

func Routing() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	http.HandleFunc("/products-static", func(w http.ResponseWriter, r *http.Request) {
		products := products.StaticWay()
		bytes, err := json.Marshal(products)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/products-json", func(w http.ResponseWriter, r *http.Request) {
		products := products.JSONBWay()
		bytes, err := json.Marshal(products)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := products.DynamicWay()
		bytes, err := json.Marshal(products)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		var product models.CreateProduct
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			w.Write([]byte(err.Error()))
		}

		id, err := products.Create(product)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(fmt.Sprintf("%d", id)))
		}
	})
}
