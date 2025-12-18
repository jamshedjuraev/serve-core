// Order - order itself. Responsible for: where is the order, whose order, status, total price
// OrderItem - what customer ordered. An order can contain 10 dishes with different quantities and prices
// OrderItemModifier - responsible for: how exactly it was prepared, extra cheese, without onions etc.
// OrderStatusHistory - order life history. Responsible for: for analitics, conflict analysis
// why did it take so long to prepare, who cancelled the order etc.

package entity

import "time"

type Order struct {
	ID           int64
	RestaurantID int64
	BranchID     int64
	TableID      int64
	Status       string // pending, in_progress, completed, cancelled
	TotalPrice   float64
	CreatedAt    time.Time
}

type OrderItem struct {
	ID         int64
	OrderID    int64
	MenuItemID int64
	Quantity   uint
	Price      float64
}

type OrderItemModifier struct {
	ID             int64
	OrderItemID    int64
	MenuModifierID int64
	PriceDelta     float64
}

type OrderStatusHistory struct {
	ID        int64
	OrderID   int64
	Status    string
	ChangedAt time.Time
	ChangedBy int64 // UserID who changed the status
}
