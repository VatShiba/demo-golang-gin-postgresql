package models

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Title       string    `json:"title" gorm:"not null" validate:"required"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
}
