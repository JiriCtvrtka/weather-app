package main

import (
	"database/sql"
	"eshop/router"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type userType struct {
	firstname, lastname, username, password, email, city, street, number, additional_info string
	zipcode, age                                                                          int64
}

type productType struct {
	id, name, description, currency string
	count                           int64
	price                           float64
}

type ordersType struct {
	id, username                string
	items                       string
	status, delivery            string
	delivery_price, total_price float64
	currency                    string
}

func main() {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort("localhost", strconv.Itoa(5432)),
		Path:     "eshop",
		User:     url.UserPassword("admin", "admin"),
		RawQuery: "sslmode=disable",
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user userType
		err := rows.Scan(&user.firstname, &user.lastname, &user.username, &user.password, &user.email, &user.city, &user.zipcode, &user.street, &user.number, &user.additional_info, &user.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var prod productType
		err := rows.Scan(&prod.id, &prod.name, &prod.description, &prod.currency, &prod.count, &prod.price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(prod)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var order ordersType
		err := rows.Scan(&order.id, &order.username, &order.items, &order.status, &order.delivery, &order.delivery_price, &order.total_price, &order.currency)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(order)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	router.Routing()

	srv := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
