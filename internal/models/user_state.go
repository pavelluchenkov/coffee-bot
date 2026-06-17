package models

import "time"

type UserState struct {
	UserID           int64
	State            string
	SelectedCategory string
	SelectedDrink    string
	UpdatedAt        time.Time
}
