package posttagmodel

import (
	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
)

const EntityName = "PostTag"

type PostTag struct {
	TagId  uuid.UUID          `json:"tag_id" gorm:"column:tag_id"`
	PostId uuid.UUID          `json:"post_id" gorm:"column:post_id"`
	Post   *common.SimplePost `json:"post"  gorm:"preload:false"`
}

func (PostTag) TableName() string {
	return "post_tags"
}
