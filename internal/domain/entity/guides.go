package entity

import (
	"time"

	"gorm.io/gorm"
)

type Guides struct {
	Model
	NameGuide          string                  `gorm:"column:name_guide;type:varchar(255);not null" json:"name_guide" validate:"required"`
	SitesID            int64                   `gorm:"column:sites_id;type:int(11);not null" json:"sites_id" validate:"required"`
	SitesName          string                  `gorm:"column:sites_name;type:varchar(255);not null" json:"sites_name" validate:"required"`
	Cover              string                  `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	State              string                  `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	MunicipalitiesID   int64                   `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id" validate:"required"`
	MunicipalitiesName string                  `gorm:"column:municipalities_name;type:varchar(255);not null" json:"municipalities_name" validate:"required"`
	DepartmentID       int64                   `gorm:"column:department_id;type:int(11);not null" json:"department_id" validate:"required"`
	DepartmentName     string                  `gorm:"column:deparment_name;type:varchar(255);not null" json:"department_name" validate:"required"`
	Description        string                  `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	Section            []SectionWihoutValidate `json:"section,omitempty"`
}

type GuidesWihoutValidate struct {
	Model
	NameGuide          string `gorm:"column:name_guide;type:varchar(255);not null" json:"name_guide" `
	SitesID            int64  `gorm:"column:sites_id;type:int(11);not null" json:"sites_id"`
	SitesName          string `gorm:"column:sites_name;type:varchar(255);not null" json:"sites_name"`
	Cover              string `gorm:"column:cover;type:varchar(255);not null" json:"cover"`
	State              string `gorm:"column:state;type:varchar(255);not null" json:"state"`
	MunicipalitiesID   int64  `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id"`
	MunicipalitiesName string `gorm:"column:municipalities_name;type:varchar(255);not null" json:"municipalities_name"`
	DepartmentID       int64  `gorm:"column:department_id;type:int(11);not null" json:"department_id"`
	DepartmentName     string `gorm:"column:deparment_name;type:varchar(255);not null" json:"department_name"`
	Description        string `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
}

func (m *GuidesWihoutValidate) TableName() string {
	return "guides"
}

func (m Guides) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Guides) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
