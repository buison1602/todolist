package storage

import "time"

type Todo struct {
	Id        int        `json:"id" gorm:"primarykey"`
	Item      string     `json:"item"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
