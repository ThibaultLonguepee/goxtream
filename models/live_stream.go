package models

import (
	"time"
)

type LiveStream struct {
	Id           int
	Number       int
	Name         string
	Icon         string
	EpgChannel   string
	AdditionDate time.Time
	IsAdult      bool
	CategoryId   int
}
