// Hierarchical structure representing menus in a restaurant
// Menu → MenuCategory → MenuItem → MenuModifier → MenuModifierOption
// Menu - restaurant's menu
// MenuCategory - category inside a menu (salad, drinks, desserts)
// MenuItem - specific dish in a category
// MenuModifier - optional features for customizing a dish (size, additional ingredients)
// MenuModifierOption - specific variant of a modifier that can change the price
// ex: roasting - rare, medium, well-done (+0), cheese - extra cheese (+2.0)

package entity

type Menu struct {
	ID           int64
	RestaurantID int64
	Name         string
	Status       string // active, inactive
}

type MenuCategory struct {
	ID     int64
	MenuID int64
	Name   string
	Status string // active, inactive
}

type MenuItem struct {
	ID             int64
	MenuCategoryID int64
	Name           string
	Description    string
	Price          float64
	ImageURL       string
	Status         string // active, inactive
}

type MenuModifier struct {
	ID         int64
	MenuItemID int64
	Name       string
	Type       string // single, multiple
	IsRequired bool
}

type MenuModifierOption struct {
	ID             int64
	MenuModifierID int64
	Name           string

	// can be 0 or positive. If pizza costs 30.0
	// and option "extra cheese" costs 5.0, total_price = 35.0
	PriceDelta float64
}
