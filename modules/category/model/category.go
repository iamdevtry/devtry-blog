package categorymodel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
)

const (
	EntityName = "category"

	TitleMaxLength = 75
	TitleMinLength = 10

	MetaTitleMaxLength = 100

	SlugMaxLength = 100
)

// Category model
type Category struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID `json:"parent_id,omitempty" gorm:"column:parent_id;"`
	Title           string     `json:"title" gorm:"column:title;"`
	MetaTitle       string     `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string     `json:"slug" gorm:"column:slug;"`
	Content         string     `json:"content" gorm:"column:content;"`
}

// Category method TableName
func (Category) TableName() string {
	return "categories"
}

// CategoryCreate model
type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID `json:"parent_id" gorm:"column:parent_id;"`
	Title           string     `json:"title" gorm:"column:title;"`
	MetaTitle       string     `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string     `json:"slug" gorm:"column:slug;"`
	Content         string     `json:"content" gorm:"column:content;"`
}

// CategoryCreate method TableName
func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

// CategoryUpdate model
type CategoryUpdate struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID `json:"parent_id" gorm:"column:parent_id;"`
	Title           string     `json:"title" gorm:"column:title;"`
	MetaTitle       string     `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string     `json:"slug" gorm:"column:slug;"`
	Content         string     `json:"content" gorm:"column:content;"`
}

// CategoryUpdate method TableName
func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

func (category *CategoryCreate) Validate() error {
	category.Title = strings.TrimSpace(category.Title)

	if len(category.Title) == 0 {
		return errors.New("title is required")
	} else if len(category.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(category.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(category.MetaTitle) > MetaTitleMaxLength {
		return fmt.Errorf("meta title is too long, max length is %d", MetaTitleMaxLength)
	}

	if len(category.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(category.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}

func (category *CategoryUpdate) Validate() error {
	category.Title = strings.TrimSpace(category.Title)

	if len(category.Title) == 0 {
		return errors.New("title is required")
	} else if len(category.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(category.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(category.MetaTitle) > MetaTitleMaxLength {
		return fmt.Errorf("meta title is too long, max length is %d", MetaTitleMaxLength)
	}

	if len(category.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(category.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}
