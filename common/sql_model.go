package common

import "time"

type SQLModel struct {
	Id     int  `json:"-" gorm:"column:id;"`
	FakeId *UID `json:"id" gorm:"-"` // FakeId is used for json response

	ParentId     int  `json:"-" gorm:"column:parent_id;"`
	FakeParentId *UID `json:"parent_id,omitempty" gorm:"-"`

	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
}

func (sql *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(sql.Id), dbType, 1)
	sql.FakeId = &uid

	if sql.ParentId != 0 {
		uparentid := NewUID(uint32(sql.ParentId), dbType, 1)
		sql.FakeParentId = &uparentid
	}

}
