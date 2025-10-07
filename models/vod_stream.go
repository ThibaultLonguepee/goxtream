package models

import "time"

type Vod struct {
	Id           int
	Number       int
	Name         string
	Icon         string
	Rating       float64
	AdditionDate time.Time
	IsAdult      bool
	Extension    string
	CategoryId   int
}
