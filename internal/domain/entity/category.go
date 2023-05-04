package entity

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	Model
	Name      string                 `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	Cover     string                 `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	State     string                 `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	SectionID int64                  `gorm:"column:section_id;type:int(11);not null" json:"section_id" validate:"required"`
	Section   *SectionWihoutValidate `gorm:"joinForeignKey:section_id;foreignKey:id;references:SectionID" json:"section,omitempty"`
}

type CategoriesWihoutValidate struct {
	Model
	Name      string                 `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	Cover     string                 `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	State     string                 `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	SectionID int64                  `gorm:"column:section_id;type:int(11);not null" json:"section_id" validate:"required"`
	Section   *SectionWihoutValidate `gorm:"joinForeignKey:section_id;foreignKey:id;references:SectionID" json:"section,omitempty"`
}

func (m CategoriesWihoutValidate) TableName() string {
	return "categories"
}

func (m Categories) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Categories) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
