package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
}

func (s *SQLModel) BeforeCreate(tx *gorm.DB) error {
	s.Id = uuid.New()
	return nil
}
