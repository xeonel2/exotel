package exotel

import (
	"errors"
)

// Auth : Defines basic auth params
type Auth struct {
	Username string
	Password string
}

func (auth *Auth) set(userName string, password string) error {
	if userName == "" || password == "" {
		return errors.New("Auth parameters missing")
	}
	auth.Username = userName
	auth.Password = password
	return nil
}
