package models

type MenuItem struct{
	ID          int
	Category    string
	Name        string
	Size        string
	Volume      int
	Price       int
	Ingredients string
	Available   bool
}
