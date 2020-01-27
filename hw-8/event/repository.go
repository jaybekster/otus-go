package calendar

import (
	"github.com/jaybekster/otus-go/hw-8/models"
)

type Repository interface {
	Create(title string) (*models.Event, error)
	Delete() (*models.Event, error)
	Edit(date string, title string) (*models.Event, error)
}