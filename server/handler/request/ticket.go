package request

import (
	"time"
)

type ReqGetTicket struct {
	Genre string `query:"genre"`
}

type ReqCreateTicket struct {
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      uint      `json:"userId"`
	GenreID     uint      `json:"genreId"`
}

type ReqUpdateTicket struct {
	ID          uint      `param:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      uint      `json:"userId"`
	GenreID     uint      `json:"genreId"`
}
