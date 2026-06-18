package repository

import (
	"coffee-bot/internal/models"
	"context"
)

func (r *Repository) CreateOrder(ctx context.Context, userID int64) (int, error) {
	var orderID int

	err := r.db.QueryRow(ctx, `
		INSERT INTO orders (user_id, status)
		VALUES ($1, 'created')
		RETURNING id
	`, userID).Scan(&orderID)

	return orderID, err
}
func (r *Repository) AddOrderItem(ctx context.Context, item models.OrderItem) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO order_items (order_id, menu_item_id, name, size, price)
		VALUES ($1, $2, $3, $4, $5)
	`, item.OrderID, item.MenuItemID, item.Name, item.Size, item.Price)

	return err
}
func (r *Repository) GetOrderTotal(ctx context.Context, orderID int) (int, error) {
	var total int

	err := r.db.QueryRow(ctx, `
		SELECT COALESCE(SUM(price),0)
		FROM order_items
		WHERE order_id = $1
	`, orderID).Scan(&total)

	return total, err
}