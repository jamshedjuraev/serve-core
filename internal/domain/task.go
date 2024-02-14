package domain

import "time"

type Task struct {
	ID          int       `json:"id,omitempty" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"not null"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"autoCreateTime"`
}
