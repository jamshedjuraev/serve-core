package entity

import "time"

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Status       string // active, inactive
	CreatedBy    int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Role struct {
	ID   int64
	Code string // root, admin, manager, kitchen
}

type UserRole struct {
	UserID       int64
	RoleID       int64
	RestaurantID int64 // nullable for root users
}
