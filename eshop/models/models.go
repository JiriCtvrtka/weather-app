package models

type UserType struct {
	firstname, lastname, username, password, email, city, street, number, additional_info string
	zipcode, age                                                                          int64
}

type ProductType struct {
	Id, Name, Description, Currency string
	Count                           int64
	Price                           float64
}

type OrdersType struct {
	id, username                string
	items                       string
	status, delivery            string
	delivery_price, total_price float64
	currency                    string
}
