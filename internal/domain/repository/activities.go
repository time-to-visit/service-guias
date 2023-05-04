package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryActivities interface {
	RegisterActivities(activities entity.Activities) (*entity.Activities, error)
	FindActivities(idSites int64) (*[]entity.Activities, error)
	FindActivitiesOne(idActivities int64) (*entity.Activities, error)
	DeleteActivities(idActivities int64) error
}

func NewRepositoryActivities(db *gorm.DB) IRepositoryActivities {
	return &repositoryActivities{
		db,
	}
}

type repositoryActivities struct {
	db *gorm.DB
}

func (r *repositoryActivities) RegisterActivities(activities entity.Activities) (*entity.Activities, error) {
	err := r.db.Create(&activities).Error
	return &activities, err
}

func (r *repositoryActivities) FindActivities(idSites int64) (*[]entity.Activities, error) {
	var activitites []entity.Activities
	err := r.db.Where("sites_id = ?", idSites).Find(&activitites).Error
	return &activitites, err
}

func (r *repositoryActivities) FindActivitiesOne(idActivities int64) (*entity.Activities, error) {
	var activity entity.Activities
	err := r.db.First(&activity, idActivities).Error
	return &activity, err
}

func (r *repositoryActivities) DeleteActivities(idActivities int64) error {
	return r.db.Delete(entity.Activities{}, idActivities).Error
}
