package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/yu-yagishita/nanpa-user/db"
	"github.com/yu-yagishita/nanpa-user/users"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service provides operations on strings.
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
	Login(username, password string) (users.User, error)
}

type service struct{}

func (service) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (service) Count(s string) int {
	return len(s)
}

func (service) Login(username, password string) (users.User, error) {
	u, err := db.GetUserByName(username)
	fmt.Println("u.FirstName: " + u.FirstName)
	if err != nil {
		return users.New(), err
	}
	if u.Password != calculatePassHash(password, u.Salt) {
		fmt.Println("u.Password: " + u.Password)
		return users.New(), ErrUnauthorized
	}

	return u, nil

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for Service.
type ServiceMiddleware func(Service) Service

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	fmt.Println(h.Sum(nil))
	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
