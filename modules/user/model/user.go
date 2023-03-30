package usermodel

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	FirstName       string     `json:"first_name" gorm:"column:first_name;"`
	MiddleName      string     `json:"middle_name" gorm:"column:middle_name;"`
	LastName        string     `json:"last_name" gorm:"column:last_name;"`
	Mobile          string     `json:"mobile" gorm:"column:mobile;"`
	Email           string     `json:"email" gorm:"column:email;"`
	Password        string     `json:"-" gorm:"column:password;"`
	Salt            string     `json:"-" gorm:"column:salt;"`
	Role            string     `json:"role" gorm:"column:role;"`
	LastLogin       *time.Time `json:"last_login" gorm:"column:last_login;"`
	Intro           string     `json:"intro" gorm:"column:intro;"`
	Profile         string     `json:"profile" gorm:"column:profile;"`
	// Avatar          string     `json:"avatar" gorm:"column:avatar;"`
	// Cover           string     `json:"cover" gorm:"column:cover;"`
}

func (u *User) GetUserId() uuid.UUID {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
