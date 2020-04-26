package main

import (
	"errors"
	"strings"
)

// Service provides operations on strings.
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
	Login(username, password string) (string, error)
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

func (service) Login(username, password string) (string, error) {
	if username == "" {
		return "", ErrEmpty
	}
	return username, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for Service.
type ServiceMiddleware func(Service) Service
