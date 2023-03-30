package common

import (
	"time"

	"github.com/google/uuid"
)

type SimplePost struct {
	SQLModel  `json:",inline"`
	AuthorId  uuid.UUID `json:"author_id" gorm:"column:author_id;"`
	Title     string    `json:"title" gorm:"column:title;"`
	MetaTitle string    `json:"meta_title" gorm:"column:meta_title;"`
	Slug      string    `json:"slug" gorm:"column:slug;"`
	Sumary    string    `json:"sumary" gorm:"column:sumary;"`
	// Thumbnail       string     `json:"thumbnail" gorm:"column:thumbnail;"`
	PublishedAt *time.Time `json:"published_at" gorm:"column:published_at;"`
}

func (SimplePost) TableName() string {
	return "posts"
}
