package users

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	u := New()
	if u.Salt == "" {
		t.Errorf("Userを作成できていない。")
	}
}

func TestValidate(t *testing.T) {
	u := New()
	err := u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Username") {
		t.Error("username が見つからないと予想されるエラー")
	}
	u.Username = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Password") {
		t.Error("password が見つからないと予想されるエラー")
	}
	u.Password = "test"
	err = u.Validate()
	if err != nil {
		t.Error(err)
	}
}
