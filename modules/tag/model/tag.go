package tagmodel

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/iamdevtry/blog/common"
)

const (
	EntityName = "post"

	TitleMaxLength = 75
	TitleMinLength = 3

	SlugMaxLength = 100
)

type Tag struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;"`
	Slug            string `json:"slug" gorm:"column:slug;"`
}

func (Tag) TableName() string {
	return "tags"
}

type TagCreate struct {
	Status    int        `json:"status" gorm:"column:status;default:1"`
	Title     string     `json:"title" gorm:"column:title;"`
	Slug      string     `json:"slug" gorm:"column:slug;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (TagCreate) TableName() string {
	return Tag{}.TableName()
}

type TagUpdate struct {
	Title string `json:"title" gorm:"column:title;"`
	Slug  string `json:"slug" gorm:"column:slug;"`
}

func (TagUpdate) TableName() string {
	return Tag{}.TableName()
}

func (tag *TagCreate) Validate() error {
	tag.Title = strings.TrimSpace(tag.Title)

	if len(tag.Title) == 0 {
		return errors.New("title is required")
	} else if len(tag.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(tag.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(tag.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(tag.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}

func (tag *TagUpdate) Validate() error {
	tag.Title = strings.TrimSpace(tag.Title)

	if len(tag.Title) == 0 {
		return errors.New("title is required")
	} else if len(tag.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(tag.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(tag.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(tag.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}
