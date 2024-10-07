package users

import (
	"fmt"

	"github.com/weather-app/eshop/db"
)

func Delete(username string) error {
	db, err := db.Connection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	user := StaticWay(username)
	if len(user) == 0 {
		return fmt.Errorf("username %s not found", username)
	}

	_, err = db.Exec("DELETE FROM users WHERE username = $1", username)

	return err
}
