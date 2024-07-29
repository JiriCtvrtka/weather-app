package db

import (
	"database/sql"
	"net"
	"net/url"
	"strconv"
)

func Connection() (*sql.DB, error) {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort("localhost", strconv.Itoa(5432)),
		Path:     "eshop",
		User:     url.UserPassword("admin", "admin"),
		RawQuery: "sslmode=disable",
	}

	db, err := sql.Open("postgres", u.String())

	return db, err
}
