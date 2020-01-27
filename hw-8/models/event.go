package models

import (
	"time"
)

type Event struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}