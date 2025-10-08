package models

import (
	"fmt"
	"time"
)

type Account struct {
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

func (srv *Server) GetUrl() string {
	return fmt.Sprintf("%v://%v:%v", srv.Protocol, srv.Url, srv.Port)
}
