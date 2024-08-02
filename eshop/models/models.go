package models

type UserType struct {
	Firstname, Lastname, Username, Password, Email, City, Street, Number, AdditionalInfo string
	Zipcode, Age                                                                         int64
}

type ProductType struct {
	Id, Name, Description, Currency string
	Count                           int64
	Price                           float64
}

type OrdersType struct {
	Id, Username              string
	Items                     string
	Status, Delivery          string
	DeliveryPrice, TotalPrice float64
	Currency                  string
}
