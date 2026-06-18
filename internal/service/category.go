package service

func CategoryName(category string) string {
	switch category {
	case "classic":
		return "Классический кофе"
	case "special":
		return "Фирменный кофе"
	case "milkshake":
		return "Милкшейки"
	case "bubble-tea":
		return "Bubble-Tea"
	default:
		return category
	}
}
func CategoryCode(name string) string {
	switch name {
	case "Классический кофе":
		return "classic"
	case "Фирменный кофе":
		return "special"
	case "Милкшейки":
		return "milkshake"
	case "Bubble-Tea":
		return "bubble-tea"
	default:
		return name
	}
}