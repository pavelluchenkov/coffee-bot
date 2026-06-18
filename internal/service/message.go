package service

import (
	"context"
	"fmt"
	"strings"

	"coffee-bot/internal/models"
)

type MessageService struct {
	menuService      *MenuService
	userStateService *UserStateService
	orderService     *OrderService
}

func NewMessageService(menuService *MenuService, userStateService *UserStateService, orderService *OrderService) *MessageService {
	return &MessageService{
		menuService:      menuService,
		userStateService: userStateService,
		orderService:     orderService,
	}
}

func (m *MessageService) ProcessMessage(ctx context.Context, userID int64, payload string) (string, error) {

	// нормализация
	payload = strings.TrimSpace(payload)

	state, err := m.userStateService.Get(ctx, userID)
	if err != nil {
		return "", err
	}

	if state == nil {
		state = &models.UserState{
			UserID: userID,
			State:  "idle",
		}
	}
	if state.State == "idle" {
		return "Нажмите кнопку 'Меню'", nil
	}

	parts := strings.Split(payload, ":")

	action := parts[0]
	value := ""
	if len(parts) > 1 {
		value = parts[1]
	}

	// -------------------------
	// START MENU
	// -------------------------
	if action == "menu" && value == "start" {

		newState := models.UserState{
			UserID: userID,
			State:  "choosing_category",
		}

		if err := m.userStateService.Save(ctx, newState); err != nil {
			return "", err
		}

		categories, err := m.menuService.GetCategories(ctx)
		if err != nil {
			return "", err
		}

		response := "Выберите категорию:\n\n"
		for _, c := range categories {
			response += CategoryName(c) + "\n"
		}

		return response, nil
	}

	// -------------------------
	// CHOOSE CATEGORY
	// -------------------------
	if state.State == "choosing_category" && action == "category" {

		category := value

		items, err := m.menuService.GetByCategory(ctx, category)
		if err != nil {
			return "", err
		}

		if len(items) == 0 {
			return "Категория не найдена", nil
		}

		newState := models.UserState{
			UserID:           userID,
			State:            "choosing_drink",
			SelectedCategory: category,
		}

		if err := m.userStateService.Save(ctx, newState); err != nil {
			return "", err
		}

		unique := make(map[string]bool)

		response := "Выберите напиток:\n\n"
		for _, item := range items {
			if !unique[item.Name] {
				unique[item.Name] = true
				response += item.Name + "\n"
			}
		}

		return response, nil
	}

	// -------------------------
	// CHOOSE DRINK
	// -------------------------
	if state.State == "choosing_drink" && action == "drink" {

		drink := value

		items, err := m.menuService.GetByName(ctx, drink)
		if err != nil {
			return "", err
		}

		if len(items) == 0 {
			return "Напиток не найден", nil
		}

		newState := models.UserState{
			UserID:           userID,
			State:            "choosing_value",
			SelectedCategory: state.SelectedCategory,
			SelectedDrink:    items[0].Name,
		}

		if err := m.userStateService.Save(ctx, newState); err != nil {
			return "", err
		}

		response := items[0].Name + "\n\n"
		for _, item := range items {
			response += fmt.Sprintf("%s — %d₽\n", item.Size, item.Price)
		}

		return response, nil
	}
	// -------------------------
	// CHOOSE SIZE
	// -------------------------
	if state.State == "choosing_value" && action == "size" {

		size := value

		if size == "" {
			return "Выберите размер", nil
		}

		// получаем все варианты напитка
		items, err := m.menuService.GetByName(ctx, state.SelectedDrink)
		if err != nil {
			return "", err
		}

		if len(items) == 0 {
			return "Напиток не найден", nil
		}

		// ищем конкретный size
		var selected models.MenuItem
		found := false

		for _, item := range items {
			if item.Size == size {
				selected = item
				found = true
				break
			}
		}

		if !found {
			return "Такого размера нет", nil
		}

		// 🔥 создаём заказ (если ещё не создан)
		orderID := state.OrderID
		if orderID == 0 {
			newOrderID, err := m.orderService.StartOrder(ctx, userID)
			if err != nil {
				return "", err
			}
			orderID = newOrderID
		}

		// 🔥 добавляем item в заказ
		err = m.orderService.AddItem(ctx, orderID, selected)
		if err != nil {
			return "", err
		}

		// сохраняем state
		newState := models.UserState{
			UserID:           userID,
			State:            "done",
			SelectedCategory: state.SelectedCategory,
			SelectedDrink:    state.SelectedDrink,
			SelectedSize:     size,
			SelectedItemID:   selected.ID,
			OrderID:          orderID,
		}

		if err := m.userStateService.Save(ctx, newState); err != nil {
			return "", err
		}

		// считаем итог
		total, _ := m.orderService.GetTotal(ctx, orderID)

		response := fmt.Sprintf(
			"Ваш выбор:\n\n%s (%s)\n%d₽\n\nИтого: %d₽",
			selected.Name,
			selected.Size,
			selected.Price,
			total,
		)

		return response, nil
	}
	// -------------------------
	// DEFAULT
	// -------------------------
	return "Нажмите кнопку меню", nil
}
