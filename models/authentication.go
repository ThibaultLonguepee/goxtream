package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
)

type Authentication struct {
	User   *User
	Server *Server
}

func ParseAuthentication(info dtos.AuthResult) (*Authentication, error) {
	user, err := ParseUser(info.UserInfo)
	if err != nil {
		return nil, err
	}

	server, err := ParseServer(info.ServerInfo)
	if err != nil {
		return nil, err
	}

	return &Authentication{
		User:   user,
		Server: server,
	}, nil
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

func ParseUser(info dtos.AuthUserInfo) (*User, error) {
	creat, err := strconv.Atoi(info.CreatedAt)
	if err != nil {
		return nil, errors.New("wrong creation date format")
	}

	exp, err := strconv.Atoi(info.ExpirationDate)
	if err != nil {
		return nil, errors.New("wrong expiration date format")
	}

	max, err := strconv.Atoi(info.MaxConnections)
	if err != nil {
		return nil, errors.New("wrong max connections format")
	}

	return &User{
		Name:           info.Username,
		Password:       info.Password,
		Trial:          info.IsTrial != "0",
		CreationDate:   time.Unix(int64(creat), 0),
		ExpirationDate: time.Unix(int64(exp), 0),
		MaxConnections: max,
		Formats:        info.AllowedOutputFormats,
	}, nil
}

type Server struct {
	Url      string
	Port     int
	Protocol string
}

func ParseServer(info dtos.AuthServerInfo) (*Server, error) {
	port, err := strconv.Atoi(info.Port)
	if err != nil {
		return nil, errors.New("invalid server port")
	}

	return &Server{
		Url:  info.Url,
		Port: port,
	}, nil
}
