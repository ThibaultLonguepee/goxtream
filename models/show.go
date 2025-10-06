package models

import "time"

type Show struct {
	Id               int
	Number           int
	Name             string
	Cover            string
	Plot             string
	Cast             []string
	Directors        []string
	Genre            string
	ReleaseDate      time.Time
	LastModification time.Time
	Rating           float64
	Backdrops        []string
	Trailer          string
	RunTime          int
	Category         int
}

type ShowDetails struct {
	Name             string
	Cover            string
	Plot             string
	Cast             []string
	Directors        []string
	Genre            string
	ReleaseDate      time.Time
	LastModification time.Time
	Rating           float64
	Backdrops        []string
	Trailer          string
	RunTime          int
	Category         int
	Seasons          []*Season
}

type Season struct {
	Id       int
	Number   int
	Name     string
	Overview string
	AirDate  time.Time
	Rating   float64
	Cover    string
	Episodes []*Episode
}

type Episode struct {
	Id           int
	Number       int
	Season       int
	Title        string
	Plot         string
	RealeaseDate time.Time
	Rating       float64
	Duration     int
	AdditionDate time.Time
	Extension    string
}
