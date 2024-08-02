package main

import (
	"fmt"
	"net/http"
	"time"

	"eshop/router"

	_ "github.com/lib/pq"
)

func main() {
	srv := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	router.Routing()
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
