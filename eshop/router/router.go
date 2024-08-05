package router

import (
	"encoding/json"

	"net/http"

	"github.com/weather-app/eshop/products"
)

func Routing() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
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
}
