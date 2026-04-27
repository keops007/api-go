package models

import "time"

type ShoppingItem struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
