package services

import (
	"github.com/shakilbd009/go-users-api/domain/users"
	"github.com/shakilbd009/go-users-api/utils/crypto_utils"
	"github.com/shakilbd009/go-users-api/utils/date_utils"
	"github.com/shakilbd009/go-users-api/utils/errors"
)

var UsersService usersServiceInterface = &usersService{}

type usersService struct{}
type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (u *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (u *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{Id: userID}
	return user.Delete()
}

func (u *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	var user users.User
	return user.FindByStatus(status)
}
