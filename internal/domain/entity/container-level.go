package entity

import (
	"time"

	"gorm.io/gorm"
)

type ContainerLevels struct {
	Model
	Type      string                 `gorm:"column:type;type:varchar(255);not null" json:"type" validate:"required"`
	Container string                 `gorm:"column:container;type:text;size:65535;not null" json:"container" validate:"required"`
	LevelsID  int64                  `gorm:"column:level_id;type:int(11);not null" json:"level_id" validate:"required"`
	Levels    *LevelsWithoutValidate `gorm:"joinForeignKey:level_id;foreignKey:id;references:LevelsID" json:"levels,omitempty"`
	State     string                 `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
}

func (m ContainerLevels) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m ContainerLevels) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
