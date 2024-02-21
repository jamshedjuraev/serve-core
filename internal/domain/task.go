package domain

import "time"

type TaskList struct {
	Page  int     `json:"page"`
	Pages int     `json:"pages"`
	Tasks []*Task `json:"tasks"`
}

type Task struct {
	ID          int       `json:"id,omitempty" gorm:"primaryKey"`
	UserID      string    `json:"user_id" gorm:"not null"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	IsDeleted *bool      `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
}
