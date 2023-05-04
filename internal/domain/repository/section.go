package repository

import (
	"service-user/internal/domain/entity"
	"service-user/internal/domain/utils"

	"gorm.io/gorm"
)

type IRepositorySection interface {
	InsertSection(section entity.Section) (*entity.Section, error)
	FindSection(filter map[string]interface{}) (*[]entity.Section, error)
	FindSectionOne(idSection int64) (*entity.Section, error)
	DeleteSection(idSection int64) error
}

func NewRepositorySection(db *gorm.DB) IRepositorySection {
	return &repositorySection{
		db,
	}
}

type repositorySection struct {
	db *gorm.DB
}

func (r *repositorySection) InsertSection(section entity.Section) (*entity.Section, error) {
	err := r.db.Create(&section).Error
	return &section, err
}

func (r *repositorySection) FindSection(filter map[string]interface{}) (*[]entity.Section, error) {
	var sections []entity.Section
	command, request := utils.GetWhere(filter)
	err := r.db.Where(command, request...).Find(&sections).Error
	return &sections, err
}

func (r *repositorySection) FindSectionOne(idSection int64) (*entity.Section, error) {
	var section entity.Section
	err := r.db.First(&section, idSection).Error
	return &section, err
}

func (r *repositorySection) DeleteSection(idSection int64) error {
	return r.db.Delete(entity.Section{}, idSection).Error
}
