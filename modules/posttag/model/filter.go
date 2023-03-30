package posttagmodel

import "github.com/google/uuid"

type Filter struct {
	TagId  uuid.UUID `json:"-" gorm:"column:tag_id"`
	PostId uuid.UUID `json:"-" gorm:"column:post_id"`
}
