package users

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Update(username string, user models.UserType) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	p := StaticWay(username)
	if len(p) == 0 {
		return fmt.Errorf("username %s not found", username)
	}

	_, err = db.Exec("UPDATE  users set (firstname, lastname, username, password, email, city, street, number, additional_info, zipcode, age) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) where username=$12",
		user.Firstname, user.Lastname, user.Username, user.Password, user.Email, user.City, user.Street, user.Number, user.AdditionalInfo, user.Zipcode, user.Age, username)

	return err
}
