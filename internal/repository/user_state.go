package repository

import (
	"coffee-bot/internal/models"
	"context"
)

func (r *Repository) GetUserState(ctx context.Context, userID int64) (*models.UserState, error){
	row := r.db.QueryRow(ctx, `SELECT user_id, state, selected_category, selected_drink, updated_at 
	FROM user_states WHERE user_id = $1`, userID)

	var state models.UserState
	err := row.Scan(&state.UserID, &state.State, &state.SelectedCategory, &state.SelectedDrink, &state.UpdatedAt)
	if err != nil{
		return nil, err
	}
	return &state, nil
}
func (r *Repository) SaveUserState(ctx context.Context, state models.UserState) error{
	_, err := r.db.Exec(ctx, `INSERT INTO user_states (user_id, state, selected_category, selected_drink) VALUES ($1, $2, $3, $4) 
	ON CONFLICT (user_id) DO UPDATE SET state = EXCLUDED.state, selected_category = EXCLUDED.selected_category, 
	selected_drink = EXCLUDED.selected_drink, updated_at = NOW()`, state.UserID, state.State, state.SelectedCategory, state.SelectedDrink)

	if err != nil{
		return err
	}
	return nil
	
}