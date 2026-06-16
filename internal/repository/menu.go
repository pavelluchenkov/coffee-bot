package repository

import (
	"coffee-bot/internal/models"
	"context"
)

func (r *Repository) GetMenuByCategory(ctx context.Context, category string)([]models.MenuItem, error){

	rows, err := r.db.Query(ctx, `SELECT id, category, name, size, volume, price, ingredients, available 
	FROM menu WHERE category = $1 AND available = true ORDER BY name, volume`, category)
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	var result []models.MenuItem

	for rows.Next(){
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Category, &item.Name, &item.Size, &item.Volume, &item.Price, &item.Ingredients, &item.Available)
		if err != nil{
			return nil, err 
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return result, nil 
}

func (r *Repository) GetMenuByName(ctx context.Context, name string)([]models.MenuItem, error){
	rows, err := r.db.Query(ctx, `SELECT id, category, name, size, volume, price, ingredients, available
	FROM menu WHERE name = $1 AND available = true ORDER BY volume`, name)
	if err != nil{
		return nil, err
	}
    
	defer rows.Close()
    
	var result []models.MenuItem

	for rows.Next(){
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Category, &item.Name, &item.Size, &item.Volume, &item.Price, &item.Ingredients, &item.Available)
		if err != nil{
			return nil, err 
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil{
		return nil, err 
	}
	return result, nil 
}
func (r *Repository) GetCategories(ctx context.Context)([]string, error){
	rows, err := r.db.Query(ctx, `SELECT DISTINCT category FROM menu WHERE available = true ORDER BY category`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next(){
		var category string
		err := rows.Scan(&category)
		if err != nil{
			return nil, err
		}
		result = append(result, category)
	}
	if err := rows.Err(); err != nil{
		return nil, err
	}
	return result, err 
}