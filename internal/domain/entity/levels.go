package entity

import (
	"time"

	"gorm.io/gorm"
)

type Levels struct {
	Model
	Title       string                    `gorm:"column:title;type:varchar(255);not null" json:"title" validate:"required"`
	State       string                    `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	Cover       string                    `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	CategoryID  int64                     `gorm:"column:category_id;type:int(11);not null" json:"category_id" validate:"required"`
	Categories  *CategoriesWihoutValidate `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryID" json:"category,omitempty"`
	Description string                    `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
}

type LevelsWithoutValidate struct {
	Model
	Title       string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	State       string `gorm:"column:state;type:varchar(255);not null" json:"state"`
	CategoryID  int64  `gorm:"column:category_id;type:int(11);not null" json:"category_id"`
	Description string `gorm:"column:description;type:varchar(255);not null" json:"description"`
}

func (m LevelsWithoutValidate) TableName() string {
	return "levels"
}

func (m Levels) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Levels) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
