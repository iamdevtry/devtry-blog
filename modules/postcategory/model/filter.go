package postcategorymodel

type Filter struct {
	PostId     int `json:"-" gorm:"column:post_id;"`
	CategoryId int `json:"-" gorm:"column:category_id;"`
}
