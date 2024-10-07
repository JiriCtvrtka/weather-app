package users

import (
	"fmt"
	"log"

	"github.com/weather-app/eshop/db"
	"github.com/weather-app/eshop/models"
)

func StaticWay(username string) []models.UserType {
	db, err := db.Connection()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		log.Fatal(err)
	}

	res := []models.UserType{}
	for rows.Next() {
		var user models.UserType
		err := rows.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Password, &user.Email, &user.City, &user.Zipcode, &user.Street, &user.Number, &user.AdditionalInfo, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		res = append(res, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
