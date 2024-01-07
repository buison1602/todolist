package storage

import "time"

type User struct {
	Id           int        `json:"id" gorm:"primarykey"`
	Email        string     `json:"email"`
	UserName     string     `json:"username"` // gorm:"unique"
	PasswordHash string     `json:"passwordHash"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}
