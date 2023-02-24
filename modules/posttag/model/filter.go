package posttagmodel

type Filter struct {
	TagId  int `json:"-" gorm:"column:tag_id"`
	PostId int `json:"-" gorm:"column:post_id"`
}
