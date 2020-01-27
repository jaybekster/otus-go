package models

type Calendar struct {
	datesMap map[string][]int64
	events map[int64]*Event
}