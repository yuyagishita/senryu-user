package users

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

var (
	// ErrNoUserInResponse はレスポンスのユーザーが違っていたらエラーを返す
	ErrNoUserInResponse = errors.New("Response has no matching user")
	// ErrMissingField は入力フォームに誤りがあるときに返す
	ErrMissingField = "Error missing %v"
)

// User はユーザー情報
type User struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"-" bson:"email"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"-" bson:"password,omitempty"`
	UserID    string `json:"id" bson:"-"`
	Salt      string `json:"-" bson:"salt"`
}

// New はユーザーを作成する
func New() User {
	u := User{}
	u.NewSalt()
	return u
}

// Validate は入力フォームのバリデーションをする
func (u *User) Validate() error {
	if u.FirstName == "" {
		return fmt.Errorf(ErrMissingField, "FirstName")
	}
	if u.LastName == "" {
		return fmt.Errorf(ErrMissingField, "LastName")
	}
	if u.Username == "" {
		return fmt.Errorf(ErrMissingField, "Username")
	}
	if u.Password == "" {
		return fmt.Errorf(ErrMissingField, "Password")
	}
	return nil
}

// NewSalt はパスワードをハッシュ化するのに使用する
func (u *User) NewSalt() {
	h := sha1.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	u.Salt = fmt.Sprintf("%x", h.Sum(nil))
}
