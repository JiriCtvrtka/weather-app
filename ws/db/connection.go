package db

import (
	"database/sql"
	"net"
	"net/url"
	"strconv"
)

func Connect() (*sql.DB, error) {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort("localhost", strconv.Itoa(5432)),
		Path:     "gotest",
		User:     url.UserPassword("gotest", "gotestpassword"),
		RawQuery: "sslmode=disable",
	}

	db, err := sql.Open("postgres", u.String())

	return db, err
}
