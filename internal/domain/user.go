package domain

import "time"

type User struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Username  string    `json:"username,omitempty" gorm:"unique;size:16"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}
