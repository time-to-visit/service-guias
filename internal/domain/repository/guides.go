package repository

import (
	"service-user/internal/domain/entity"
	"service-user/internal/domain/utils"

	"gorm.io/gorm"
)

type IRepositoryGuides interface {
	InsertGuides(guides entity.Guides) (*entity.Guides, error)
	FindGuides(filter map[string]interface{}) (*[]entity.Guides, error)
	FindGuidesOne(idGuide int64) (*entity.Guides, error)
	DeleteGuides(idGuides int64) error
}

func NewRepositoryGuides(db *gorm.DB) IRepositoryGuides {
	return &repositoryGuides{
		db,
	}
}

type repositoryGuides struct {
	db *gorm.DB
}

func (r *repositoryGuides) InsertGuides(guides entity.Guides) (*entity.Guides, error) {
	err := r.db.Create(&guides).Error
	return &guides, err
}

func (r *repositoryGuides) FindGuides(filter map[string]interface{}) (*[]entity.Guides, error) {
	var guides []entity.Guides
	command, request := utils.GetWhere(filter)
	err := r.db.Where(command, request...).Find(&guides).Error
	return &guides, err

}

func (r *repositoryGuides) FindGuidesOne(idGuide int64) (*entity.Guides, error) {
	var guide entity.Guides
	var sections []entity.SectionWihoutValidate
	err := r.db.First(&guide, idGuide).Error
	r.db.Preload("Categories").Find(&sections, " guides_id =?", idGuide)
	guide.Section = sections
	return &guide, err
}

func (r *repositoryGuides) DeleteGuides(idGuides int64) error {
	return r.db.Delete(entity.Guides{}, idGuides).Error
}
