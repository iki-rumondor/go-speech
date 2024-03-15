package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	TeacherID uint   `gorm:"not_null"`
	Name      string `gorm:"not_null;size:32"`
	Code      string `gorm:"not_null;unique;size:8"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Teacher   *Teacher
}

func (m *Class) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
