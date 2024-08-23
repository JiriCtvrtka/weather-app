package router

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/weather-app/eshop/models"
	"github.com/weather-app/eshop/products"
)

func prefailedCondition(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusPreconditionFailed)
	w.Write([]byte(err.Error()))
}

func Routing() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	http.HandleFunc("/products-static", func(w http.ResponseWriter, r *http.Request) {
		products := products.StaticWay("")
		bytes, err := json.Marshal(products)

		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/products-json", func(w http.ResponseWriter, r *http.Request) {
		products := products.JSONBWay()
		bytes, err := json.Marshal(products)

		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := products.DynamicWay()
		bytes, err := json.Marshal(products)

		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		var product models.ProductCore
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		id, err := products.Create(product)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Write([]byte(fmt.Sprintf("%d", id)))
	})

	http.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		switch r.Method {
		case "GET":
			products := products.StaticWay(id)
			bytes, err := json.Marshal(products)

			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)
		case "DELETE":
			err := products.Delete(id)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Write([]byte("OK"))
		case "PUT", "PATCH":
			var product models.ProductCore
			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			err = products.Update(id, product)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Write([]byte("OK"))
		}
	})
}
