package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

type userType struct {
	firstname, lastname, address, age string
}

func main() {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort("localhost", strconv.Itoa(5432)),
		Path:     "eshop",
		User:     url.UserPassword("user", "password"),
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
		err := rows.Scan(&user.firstname, &user.lastname, &user.address, &user.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
