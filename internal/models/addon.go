package models

type Addon struct {
	ID        int
	Type      string
	Name      string
	Price     int
	Available bool
}