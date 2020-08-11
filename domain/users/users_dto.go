package users

import (
	"strings"

	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

const (
	StatusActive = "active"
)

type (
	User struct {
		Id          int64  `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		DateCreated string `json:"date_created"`
		Status      string `json:"status"`
		Password    string `json:"password"`
	}
	Users []User
)

func (user *User) Validate() rest_errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
