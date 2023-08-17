package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	Model
	Image            string               `gorm:"column:image;type:varchar(255);not null" json:"image" validate:"required"`
	Video            string               `gorm:"column:video;type:varchar(255);not null" json:"video" validate:"required"`
	Description      string               `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	SitesId          int64                `gorm:"column:department_id;type:int(11);not null" json:"site_id"`
	SitesName        string               `gorm:"column:sites_name;type:varchar(255);not null" json:"sites_name" validate:"required"`
	MunicipalitiesID int64                `gorm:"column:department_id;type:int(11);not null" json:"municipalities_id"`
	DepartmentID     int64                `gorm:"column:department_id;type:int(11);not null" json:"department_id"`
	State            string               `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	GuideId          int64                `gorm:"column:guide_id;type:int(11);not null" json:"guide_id" validate:"required"`
	Guide            GuidesWihoutValidate `gorm:"joinForeignKey:guide_id;foreignKey:id;references:GuideId" json:"guide,omitempty"`
}

type ActivitiesWihoutValidate struct {
	Model
	Image            string `gorm:"column:image;type:varchar(255);not null" json:"image" validate:"required"`
	Video            string `gorm:"column:video;type:varchar(255);not null" json:"video" validate:"required"`
	Description      string `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	SitesId          int64  `gorm:"column:department_id;type:int(11);not null" json:"site_id"`
	SitesName        string `gorm:"column:sites_name;type:varchar(255);not null" json:"sites_name" validate:"required"`
	MunicipalitiesID int64  `gorm:"column:department_id;type:int(11);not null" json:"municipalities_id"`
	DepartmentID     int64  `gorm:"column:department_id;type:int(11);not null" json:"department_id"`
	State            string `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	GuideId          int64  `gorm:"column:guide_id;type:int(11);not null" json:"guide_id" validate:"required"`
}

func (m *ActivitiesWihoutValidate) TableName() string {
	return "activities"
}

func (m Activities) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Activities) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
