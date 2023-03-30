package usermodel

import (
	"errors"
	"fmt"
	"time"

	"github.com/iamdevtry/blog/common"
)

const (
	MinPasswordLength   = 6
	MaxFirstNameLength  = 50
	MaxMiddleNameLength = 50
	MaxLastNameLength   = 50
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	FirstName       string     `json:"first_name" gorm:"column:first_name;"`
	MiddleName      string     `json:"middle_name" gorm:"column:middle_name;"`
	LastName        string     `json:"last_name" gorm:"column:last_name;"`
	Mobile          string     `json:"mobile" gorm:"column:mobile;"`
	Email           string     `json:"email" gorm:"column:email;"`
	Password        string     `json:"password" gorm:"column:password;"`
	Salt            string     `json:"-" gorm:"column:salt;"`
	Role            string     `json:"-" gorm:"column:role;"`
	CreatedAt       *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	Status          int        `json:"status" gorm:"column:status;default:1"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Validate() error {
	if len(u.Mobile) == 0 {
		return errors.New("mobile is required")
	} else {
		if err := common.ValidPhoneNumber(u.Mobile); err != nil {
			return err
		}
	}

	if len(u.Email) == 0 {
		return errors.New("email is required")
	} else {
		if err := common.ValidMailAddress(u.Email); err != nil {
			return err
		}
	}

	if len(u.Password) == 0 {
		return errors.New("password is required")
	} else {
		if len(u.Password) < 6 {
			return fmt.Errorf("password is too short, minimum %d characters", MinPasswordLength)
		}
	}

	if len(u.FirstName) > MaxFirstNameLength {
		return fmt.Errorf("first name is too long, maximum %d characters", MaxFirstNameLength)
	}

	if len(u.MiddleName) > MaxMiddleNameLength {
		return fmt.Errorf("middle name is too long, maximum %d characters", MaxMiddleNameLength)
	}

	if len(u.LastName) > MaxLastNameLength {
		return fmt.Errorf("last name is too long, maximum %d characters", MaxLastNameLength)
	}

	return nil
}
