package entity

import "time"

type Kitchen struct {
	ID         int64
	OrderID    int64
	StartedAt  time.Time
	FinishedAt time.Time
}
