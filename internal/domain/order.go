package domain

// Order описывает заказ на уровне предметной области.
// Здесь нет знаний о том, как заказ хранится в базе данных.
type Order struct {
	ID       int
	Customer string
	Products []string
	Total    float64
	Status   string
}
