// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-user/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryQuestion is an autogenerated mock type for the IRepositoryQuestion type
type IRepositoryQuestion struct {
	mock.Mock
}

// DeleteQuestion provides a mock function with given fields: idQuestion
func (_m *IRepositoryQuestion) DeleteQuestion(idQuestion int64) error {
	ret := _m.Called(idQuestion)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(idQuestion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindQuestionByObjective provides a mock function with given fields: idObjective
func (_m *IRepositoryQuestion) FindQuestionByObjective(idObjective int64) (*[]entity.Question, error) {
	ret := _m.Called(idObjective)

	var r0 *[]entity.Question
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*[]entity.Question, error)); ok {
		return rf(idObjective)
	}
	if rf, ok := ret.Get(0).(func(int64) *[]entity.Question); ok {
		r0 = rf(idObjective)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Question)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idObjective)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertQuestion provides a mock function with given fields: question
func (_m *IRepositoryQuestion) InsertQuestion(question entity.Question) (*entity.Question, error) {
	ret := _m.Called(question)

	var r0 *entity.Question
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Question) (*entity.Question, error)); ok {
		return rf(question)
	}
	if rf, ok := ret.Get(0).(func(entity.Question) *entity.Question); ok {
		r0 = rf(question)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Question)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Question) error); ok {
		r1 = rf(question)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryQuestion interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryQuestion creates a new instance of IRepositoryQuestion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryQuestion(t mockConstructorTestingTNewIRepositoryQuestion) *IRepositoryQuestion {
	mock := &IRepositoryQuestion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
