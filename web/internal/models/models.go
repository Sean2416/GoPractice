package models

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Reservation holds reservation data
type Payment struct {
	OrderNo      string
	Desc         string
	Amount       float32
	Currency     string
	Email        string
	Phone        string
	CustomerName string
}
