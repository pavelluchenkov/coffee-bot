package handler

import (
	"coffee-bot/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type MenuHandler struct {
	menuService *service.MenuService
}

func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (m *MenuHandler) GetCategories(w http.ResponseWriter, r *http.Request){
	
	categories, err := m.menuService.GetCategories(r.Context()) 
	if err != nil{
		http.Error(w, "failed to get categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") 

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
func (m *MenuHandler) GetByCategory(w http.ResponseWriter, r *http.Request){
	
	parts := strings.Split(r.URL.Path, "/")
	category := parts[2]

	items, err := m.menuService.GetByCategory(r.Context(), category)
	if err != nil{
		http.Error(w, "failed to get menu in current category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(items); err != nil{
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}