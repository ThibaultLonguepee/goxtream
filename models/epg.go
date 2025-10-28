package models

import "time"

type Epg struct {
	Id          int
	Title       string
	Description string
	Start       time.Time
	End         time.Time
	Language    string
}
