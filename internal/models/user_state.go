package models

import (
	"time"
)

type UserState struct {
	UserID           int64
	State            string
	SelectedCategory string
	SelectedDrink    string
	SelectedSize     string
	SelectedItemID   int
	OrderID          int
	UpdatedAt        time.Time
}
