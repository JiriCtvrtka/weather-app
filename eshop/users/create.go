package users

import (
	"fmt"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func Create(item models.UserType) (string, error) {
	db, err := db.Connection()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()

	var username string
	err = db.QueryRow(
		"INSERT INTO users (firstname, lastname, username, password, email, city, street, number, additional_info, zipcode, age) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning username",
		item.Firstname, item.Lastname, item.Username, item.Password, item.Email, item.City, item.Street, item.Number, item.AdditionalInfo, item.Zipcode, item.Age).Scan(&username)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return username, nil
}
