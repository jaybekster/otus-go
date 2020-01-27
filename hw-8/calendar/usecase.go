package calendar

import (
	"github.com/jaybekster/otus-go/hw-8/models"
)

type Usecase interface {
	AddEvent(date string, title string) (*models.Event, error)
	RemoveEvent(id int64) (*models.Event, error)
	EditEvent(date string, title string) (*models.Event, error)
	ListEvents() (map[string][]*models.Event)
}