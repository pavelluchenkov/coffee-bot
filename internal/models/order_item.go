package models

type OrderItem struct {
	ID          int
	OrderID     int
	MenuItemID  int
	Name        string
	Size        string
	Price       int
}