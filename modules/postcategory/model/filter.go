package postcategorymodel

import "github.com/google/uuid"

type Filter struct {
	PostId     uuid.UUID `json:"-" gorm:"column:post_id;"`
	CategoryId uuid.UUID `json:"-" gorm:"column:category_id;"`
}
