package entity

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	Model
	Type        string                    `gorm:"column:type;type:varchar(255);not null" json:"type" validate:"required"`
	Request     string                    `gorm:"column:request;type:varchar(255);not null" json:"request" validate:"required"`
	IsTrue      string                    `gorm:"column:is_true;type:varchar(255);not null" json:"is_true" validate:"required"`
	ObjectiveID int64                     `gorm:"column:objective_id;type:int(11);not null" json:"objective_id" validate:"required"`
	Objectives  *ObjectivesWihoutValidate `gorm:"joinForeignKey:objective_id;foreignKey:id;references:ObjectiveID" json:"objectives,omitempty"`
	State       string                    `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
}

func (m Question) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Question) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
