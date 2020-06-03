package api

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
	//ErrUnauthorized エラーになった際に返す値
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service はAPIで提供している機能をまとめたインターフェース
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
	Login(username, password string) (users.User, error)
	Register(username, email, password string) (users.User, error)
}

// NewFixedService はfixedService{}を返す
func NewFixedService() Service {
	return &fixedService{}
}

type fixedService struct{}

func (s *fixedService) Uppercase(str string) (string, error) {
	if str == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(str), nil
}

func (s *fixedService) Count(str string) int {
	return len(str)
}

func (s *fixedService) Login(username, password string) (users.User, error) {
	u, err := db.GetUserByName(username)
	if err != nil {
		return users.New(), err
	}
	if u.Password != calculatePassHash(password, u.Salt) {
		return users.New(), ErrUnauthorized
	}

	return u, nil

}

func (s *fixedService) Register(username, email, password string) (users.User, error) {
	u := users.New()
	u.Username = username
	u.Email = email
	u.Password = calculatePassHash(password, u.Salt)
	err := db.CreateUser(&u)
	return u.UserID, err
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
