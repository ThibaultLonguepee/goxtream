package mappers

import (
	"errors"
	"strconv"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
	"github.com/thibaultlonguepee/goxtream/models"
)

func ParseAccount(info dtos.AuthResult) (*models.Account, error) {
	user, err := ParseUser(info.UserInfo)
	if err != nil {
		return nil, err
	}

	server, err := ParseServer(info.ServerInfo)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		User:   user,
		Server: server,
	}, nil
}

func ParseUser(info dtos.AuthUserInfo) (*models.User, error) {
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

	return &models.User{
		Name:           info.Username,
		Password:       info.Password,
		Trial:          info.IsTrial != "0",
		CreationDate:   time.Unix(int64(creat), 0),
		ExpirationDate: time.Unix(int64(exp), 0),
		MaxConnections: max,
		Formats:        info.AllowedOutputFormats,
	}, nil
}

func ParseServer(info dtos.AuthServerInfo) (*models.Server, error) {
	port, err := strconv.Atoi(info.Port)
	if err != nil {
		return nil, errors.New("invalid server port")
	}

	return &models.Server{
		Url:      info.Url,
		Port:     port,
		Protocol: info.ServerProtocol,
	}, nil
}
