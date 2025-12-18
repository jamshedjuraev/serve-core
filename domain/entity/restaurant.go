package entity

import "time"

type Restaurant struct {
	ID        int64
	Name      string
	Status    string // active, inactive
	CreatedAt time.Time
}

type Branch struct {
	ID           int64
	RestaurantID int64
	Name         string
	Address      string
	IsMain       bool
}

type Table struct {
	ID       int64
	BranchID int64
	Number   string
	Capacity int8
	Status   string // available, occupied, reserved
}

type QRCode struct {
	ID        int64
	TableID   int64
	Token     string
	ExpiresAt time.Time
}

