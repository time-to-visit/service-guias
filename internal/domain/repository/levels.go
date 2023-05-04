package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryLevels interface {
	InsertLevels(levels entity.Levels) (*entity.Levels, error)
	FindLevelsByCategory(idCategory int64) (*[]entity.Levels, error)
	FindLevelsOne(idLevels int64) (*entity.Levels, error)
	DeleteLevels(idLevels int64) error
}

func NewRepositoryLevels(db *gorm.DB) IRepositoryLevels {
	return &repositoryLevels{
		db,
	}
}

type repositoryLevels struct {
	db *gorm.DB
}

func (r *repositoryLevels) InsertLevels(levels entity.Levels) (*entity.Levels, error) {
	err := r.db.Create(&levels).Error
	return &levels, err
}

func (r *repositoryLevels) FindLevelsByCategory(idCategory int64) (*[]entity.Levels, error) {
	var levels []entity.Levels
	err := r.db.Where("category_id = ?", idCategory).Find(&levels).Error
	return &levels, err
}

func (r *repositoryLevels) FindLevelsOne(idLevels int64) (*entity.Levels, error) {
	var level entity.Levels
	err := r.db.First(&level, idLevels).Error
	return &level, err
}

func (r *repositoryLevels) DeleteLevels(idLevels int64) error {
	return r.db.Delete(entity.Levels{}, idLevels).Error
}
