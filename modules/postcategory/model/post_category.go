package postcategorymodel

import (
	"time"

	"github.com/iamdevtry/blog/common"
)

const EntityName = "PostCategory"

type PostCategory struct {
	CategoryId int                `json:"category_id" gorm:"column:category_id;"`
	PostId     int                `json:"post_id" gorm:"column:post_id;"`
	CreatedAt  *time.Time         `json:"created_at" gorm:"column:created_at"`
	Post       *common.SimplePost `json:"post"  gorm:"preload:false"`
}

func (PostCategory) TableName() string {
	return "post_categories"
}

// func (postCategory *PostCategory) Validate() error {

// }
