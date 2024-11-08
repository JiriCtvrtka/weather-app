package router

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/weather-app/eshop/models"
	"github.com/weather-app/eshop/orders"
	"github.com/weather-app/eshop/products"
	"github.com/weather-app/eshop/users"
	"github.com/weather-app/eshop/utils"
)

func prefailedCondition(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusPreconditionFailed)
	w.Write([]byte(err.Error()))
}

func Routing() {
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

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders := utils.DynamicWay("orders")
		bytes, err := json.Marshal(orders)

		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		var order models.OrdersCore
		err := json.NewDecoder(r.Body).Decode(&order)
		fmt.Println(order)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		id, err := orders.Create(order)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Write([]byte(fmt.Sprintf("%d", id)))
	})

	http.HandleFunc("/order/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		switch r.Method {
		case "GET":
			orders := orders.StaticWay(id)
			bytes, err := json.Marshal(orders)

			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)

		case "DELETE":
			err := orders.Delete(id)
			if err != nil {
				prefailedCondition(w, err)
				return
			}
			w.Write([]byte("OK"))

		case "PUT", "PATCH":
			var order models.OrdersCore
			err := json.NewDecoder(r.Body).Decode(&order)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			err = orders.Update(id, order)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Write([]byte("OK"))
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := utils.DynamicWay("users")
		bytes, err := json.Marshal(users)

		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var user models.UserType
		err := json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		username, err := users.Create(user)
		if err != nil {
			prefailedCondition(w, err)
			return
		}

		w.Write([]byte(fmt.Sprintf("%s", username)))
	})

	http.HandleFunc("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")
		switch r.Method {
		case "GET":
			user := users.StaticWay(username)
			bytes, err := json.Marshal(user)

			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)

		case "DELETE":
			err := users.Delete(username)
			if err != nil {
				prefailedCondition(w, err)
				return
			}
			w.Write([]byte("OK"))

		case "PUT", "PATCH":
			var user models.UserType
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			err = users.Update(username, user)
			if err != nil {
				prefailedCondition(w, err)
				return
			}

			w.Write([]byte("OK"))
		}
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "build/index.html") })

	fs := http.FileServer(http.Dir("build"))
	http.Handle("/react/", http.StripPrefix("/react/", fs))

}
