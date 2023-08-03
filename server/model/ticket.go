package model

import (
	"time"
)

type Ticket struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      uint      `json:"userId"`
	GenreID     uint      `json:"genreId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	User        User      `gorm:"foreignKey:UserID"`
	Genre       Genre     `gorm:"foreignKey:GenreID"`
}
