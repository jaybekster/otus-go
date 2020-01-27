package calendar

type Usecase interface {
	AddEvent()
	RemoveEvent()
	EditEvent()
	ListEvents()
}