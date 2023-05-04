package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryContainerLevels interface {
	InsertContainerLevels(containerLevels entity.ContainerLevels) (*entity.ContainerLevels, error)
	FindContainerLevels(idLevels int64) (*[]entity.ContainerLevels, error)
	FindContainerLevelsOne(idContainerLevels int64) (*entity.ContainerLevels, error)
	DeleteContainerLevels(idContainerLevels int64) error
}

func NewRepositoryContainerLevels(db *gorm.DB) IRepositoryContainerLevels {
	return &repositoryContainerLevels{
		db,
	}
}

type repositoryContainerLevels struct {
	db *gorm.DB
}

func (r *repositoryContainerLevels) InsertContainerLevels(containerLevels entity.ContainerLevels) (*entity.ContainerLevels, error) {
	err := r.db.Create(&containerLevels).Error
	return &containerLevels, err
}

func (r *repositoryContainerLevels) FindContainerLevels(idLevels int64) (*[]entity.ContainerLevels, error) {
	var containerLevels []entity.ContainerLevels
	err := r.db.Where("level_id = ?", idLevels).Find(&containerLevels).Error
	return &containerLevels, err
}

func (r *repositoryContainerLevels) FindContainerLevelsOne(idContainerLevels int64) (*entity.ContainerLevels, error) {
	var containerLevel entity.ContainerLevels
	err := r.db.First(&containerLevel, idContainerLevels).Error
	return &containerLevel, err
}

func (r *repositoryContainerLevels) DeleteContainerLevels(idContainerLevels int64) error {
	return r.db.Delete(entity.ContainerLevels{}, idContainerLevels).Error
}
