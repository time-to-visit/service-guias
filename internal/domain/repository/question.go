package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryQuestion interface {
	InsertQuestion(question entity.Question) (*entity.Question, error)
	DeleteQuestion(idQuestion int64) error
	FindQuestionByObjective(idObjective int64) (*[]entity.Question, error)
}

func NewRepositoryQuestion(db *gorm.DB) IRepositoryQuestion {
	return &repositoryQuestion{
		db,
	}
}

type repositoryQuestion struct {
	db *gorm.DB
}

func (r *repositoryQuestion) InsertQuestion(question entity.Question) (*entity.Question, error) {
	err := r.db.Create(&question).Error
	return &question, err
}

func (r *repositoryQuestion) DeleteQuestion(idQuestion int64) error {
	return r.db.Delete(entity.Question{}).Error
}

func (r *repositoryQuestion) FindQuestionByObjective(idObjective int64) (*[]entity.Question, error) {
	var questions []entity.Question
	err := r.db.Where("objective_id =?", idObjective).Find(&questions).Error
	return &questions, err
}
