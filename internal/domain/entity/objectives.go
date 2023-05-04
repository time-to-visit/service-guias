package entity

import (
	"time"

	"gorm.io/gorm"
)

type Objectives struct {
	Model
	Question string                 `gorm:"column:question;type:varchar(255);not null" json:"question" validate:"required"`
	State    string                 `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	Answer   string                 `gorm:"column:answer;type:varchar(255);not null" json:"answer" validate:"required"`
	LevelsId int64                  `gorm:"column:levels_id;type:int(11);not null" json:"levels_id" validate:"required"`
	Levels   *LevelsWithoutValidate `gorm:"joinForeignKey:levels_id;foreignKey:id;references:LevelsId" json:"levels,omitempty"`
	Ques     *[]Question            `gorm:"foreignKey:ObjectiveID;references:ID"  json:"questions,omitempty"`
}

type ObjectivesWihoutValidate struct {
	Model
	Question string `gorm:"column:question;type:varchar(255);not null" json:"question" validate:"required"`
	State    string `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	Answer   string `gorm:"column:answer;type:varchar(255);not null" json:"answer" validate:"required"`
	LevelsId int64  `gorm:"column:levels_id;type:int(11);not null" json:"levels_id" validate:"required"`
}

func (m *ObjectivesWihoutValidate) TableName() string {
	return "objectives"
}

func (m Objectives) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Objectives) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
