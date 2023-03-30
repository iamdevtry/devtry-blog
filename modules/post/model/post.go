package postmodel

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
)

const (
	EntityName = "post"

	TitleMaxLength = 75
	TitleMinLength = 30

	MetaTitleMaxLength = 100

	SlugMaxLength = 100
)

type Post struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID    `json:"parent_id,omitempty" gorm:"column:parent_id;"`
	AuthorId        uuid.UUID     `json:"author_id" gorm:"column:author_id;"`
	Title           string        `json:"title" gorm:"column:title;"`
	MetaTitle       string        `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string        `json:"slug" gorm:"column:slug;"`
	Sumary          string        `json:"sumary" gorm:"column:sumary;"`
	Content         string        `json:"content" gorm:"column:content;"`
	Thumbnail       *common.Image `json:"thumbnail" gorm:"column:thumbnail;"`
	PublishedAt     *time.Time    `json:"published_at" gorm:"column:published_at;"`
}

func (Post) TableName() string {
	return "posts"
}

type PostCreate struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID    `json:"parent_id" gorm:"column:parent_id;"`
	AuthorId        uuid.UUID     `json:"author_id" gorm:"column:author_id;"`
	Title           string        `json:"title" gorm:"column:title;"`
	MetaTitle       string        `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string        `json:"slug" gorm:"column:slug;"`
	Sumary          string        `json:"sumary" gorm:"column:sumary;"`
	Content         string        `json:"content" gorm:"column:content;"`
	Thumbnail       *common.Image `json:"thumbnail" gorm:"column:thumbnail;"`
}

func (PostCreate) TableName() string {
	return Post{}.TableName()
}

type PostUpdate struct {
	common.SQLModel `json:",inline"`
	ParentId        *uuid.UUID    `json:"parent_id" gorm:"column:parent_id;"`
	AuthorId        uuid.UUID     `json:"author_id" gorm:"column:author_id;"`
	Title           string        `json:"title" gorm:"column:title;"`
	MetaTitle       string        `json:"meta_title" gorm:"column:meta_title;"`
	Slug            string        `json:"slug" gorm:"column:slug;"`
	Sumary          string        `json:"sumary" gorm:"column:sumary;"`
	Content         string        `json:"content" gorm:"column:content;"`
	Thumbnail       *common.Image `json:"thumbnail" gorm:"column:thumbnail;"`
}

func (PostUpdate) TableName() string {
	return Post{}.TableName()
}

func (post *PostCreate) Validate() error {
	post.Title = strings.TrimSpace(post.Title)

	if len(post.Title) == 0 {
		return errors.New("title is required")
	} else if len(post.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(post.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(post.MetaTitle) > MetaTitleMaxLength {
		return fmt.Errorf("meta title is too long, max length is %d", MetaTitleMaxLength)
	}

	if len(post.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(post.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}

func (post *PostUpdate) Validate() error {
	post.Title = strings.TrimSpace(post.Title)

	if len(post.Title) == 0 {
		return errors.New("title is required")
	} else if len(post.Title) > TitleMaxLength {
		return fmt.Errorf("title is too long, max length is %d", TitleMaxLength)
	} else if len(post.Title) < TitleMinLength {
		return fmt.Errorf("title is too short, min length is %d", TitleMinLength)
	}

	if len(post.MetaTitle) > MetaTitleMaxLength {
		return fmt.Errorf("meta title is too long, max length is %d", MetaTitleMaxLength)
	}

	if len(post.Slug) == 0 {
		return errors.New("slug is required")
	} else if len(post.Slug) > SlugMaxLength {
		return fmt.Errorf("slug is too long, max length is %d", SlugMaxLength)
	}

	return nil
}
