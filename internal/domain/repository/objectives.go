package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryObjectives interface {
	InsertObjective(objective entity.Objectives) (*entity.Objectives, error)
	DeleteObjective(idObjective int64) error
	FindObjectiveByLevel(idLevels int64) (*[]entity.Objectives, error)
	FindObjectiveOne(idObjective int64) (*entity.Objectives, error)
}

func NewRepositoryObjectives(db *gorm.DB) IRepositoryObjectives {
	return &repositoryObjectives{
		db,
	}
}

type repositoryObjectives struct {
	db *gorm.DB
}

func (r *repositoryObjectives) InsertObjective(objective entity.Objectives) (*entity.Objectives, error) {
	err := r.db.Create(&objective).Error
	return &objective, err
}

func (r *repositoryObjectives) DeleteObjective(idObjective int64) error {
	return r.db.Delete(entity.Objectives{}, idObjective).Error
}

func (r *repositoryObjectives) FindObjectiveByLevel(idLevels int64) (*[]entity.Objectives, error) {
	var objectives []entity.Objectives
	err := r.db.Preload("Ques").Where("levels_id = ?", idLevels).Find(&objectives).Error
	return &objectives, err
}

func (r *repositoryObjectives) FindObjectiveOne(idObjective int64) (*entity.Objectives, error) {
	var objective entity.Objectives
	err := r.db.First(&objective, idObjective).Error
	return &objective, err
}
