package models

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
