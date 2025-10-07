package models

import "time"

type Vod struct {
	Id           int
	Number       int
	Name         string
	Thumbnail    string
	Rating       float64
	AdditionDate time.Time
	IsAdult      bool
	Extension    string
	Category     int
}

type VodDetails struct {
	Id           int
	Name         string
	OriginalName string
	Plot         string
	Genre        string
	Cast         []string
	Directors    []string
	Duration     int
	RealeaseDate time.Time
	Country      string
	Age          string
	MpaaRating   string
	Rating       float64
	RatingCount  int
	AdditionDate time.Time
	Cover        string
	Backdrops    []string
	Bitrate      int
	Extension    string
	Category     int
}
