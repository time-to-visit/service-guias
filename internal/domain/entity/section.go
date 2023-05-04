package entity

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	Model
	Name       string                     `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	Cover      string                     `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	State      string                     `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	GuideId    int64                      `gorm:"column:guides_id;type:int(11);not null" json:"guides_id" validate:"required"`
	Guides     *GuidesWihoutValidate      `gorm:"joinForeignKey:guides_id;foreignKey:id;references:GuideId" json:"guides,omitempty"`
	Categories []CategoriesWihoutValidate `gorm:"foreignKey:SectionID;references:ID"  json:"categories,omitempty"`
}

type SectionWihoutValidate struct {
	Model
	Name       string                     `gorm:"column:name;type:varchar(255);not null" json:"name" `
	Cover      string                     `gorm:"column:cover;type:varchar(255);not null" json:"cover" `
	State      string                     `gorm:"column:state;type:varchar(255);not null" json:"state"`
	GuideId    int64                      `gorm:"column:guides_id;type:int(11);not null" json:"guides_id"`
	Categories []CategoriesWihoutValidate `gorm:"foreignKey:SectionID;references:ID"  json:"categories,omitempty"`
}

func (m *SectionWihoutValidate) TableName() string {
	return "sections"
}

func (m Section) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Section) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
