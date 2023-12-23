package models

import "time"

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Progress    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
