package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryCategories interface {
	InsertCategory(categories entity.Categories) (*entity.Categories, error)
	FindCategory(idSection int64) (*[]entity.Categories, error)
	FindCategoryOne(idCategory int64) (*entity.Categories, error)
	DeleteCategory(idCategory int64) error
}

func NewRepositoryCategories(db *gorm.DB) IRepositoryCategories {
	return &repositoryCategories{
		db,
	}
}

type repositoryCategories struct {
	db *gorm.DB
}

func (r *repositoryCategories) InsertCategory(categories entity.Categories) (*entity.Categories, error) {
	err := r.db.Create(&categories).Error
	return &categories, err
}

func (r *repositoryCategories) FindCategory(idSection int64) (*[]entity.Categories, error) {
	var categories []entity.Categories
	err := r.db.Where("section_id = ?", idSection).Find(&categories).Error
	return &categories, err
}

func (r *repositoryCategories) FindCategoryOne(idCategory int64) (*entity.Categories, error) {
	var category entity.Categories
	err := r.db.First(&category, idCategory).Error
	return &category, err
}

func (r *repositoryCategories) DeleteCategory(idCategory int64) error {
	return r.db.Delete(entity.Categories{}, idCategory).Error
}
