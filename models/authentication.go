package models

import (
	"time"
)

type Authentication struct {
	User   *User
	Server *Server
}

type User struct {
	Name           string
	Password       string
	Trial          bool
	CreationDate   time.Time
	ExpirationDate time.Time
	MaxConnections int
	Formats        []string
}

type Server struct {
	Url      string
	Port     int
	Protocol string
}
