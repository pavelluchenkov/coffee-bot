package models

type Order struct {

	ID         int
	UserID     int64
	Status     string
	TotalPrice int

}