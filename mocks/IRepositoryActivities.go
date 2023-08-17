// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-user/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryActivities is an autogenerated mock type for the IRepositoryActivities type
type IRepositoryActivities struct {
	mock.Mock
}

// DeleteActivities provides a mock function with given fields: idActivities
func (_m *IRepositoryActivities) DeleteActivities(idActivities int64) error {
	ret := _m.Called(idActivities)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(idActivities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindActivities provides a mock function with given fields: idSites
func (_m *IRepositoryActivities) FindActivities(idSites int64) (*[]entity.Activities, error) {
	ret := _m.Called(idSites)

	var r0 *[]entity.Activities
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*[]entity.Activities, error)); ok {
		return rf(idSites)
	}
	if rf, ok := ret.Get(0).(func(int64) *[]entity.Activities); ok {
		r0 = rf(idSites)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Activities)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idSites)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActivitiesOne provides a mock function with given fields: idActivities
func (_m *IRepositoryActivities) FindActivitiesOne(idActivities int64) (*entity.Activities, error) {
	ret := _m.Called(idActivities)

	var r0 *entity.Activities
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*entity.Activities, error)); ok {
		return rf(idActivities)
	}
	if rf, ok := ret.Get(0).(func(int64) *entity.Activities); ok {
		r0 = rf(idActivities)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Activities)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idActivities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterActivities provides a mock function with given fields: activities
func (_m *IRepositoryActivities) RegisterActivities(activities entity.Activities) (*entity.Activities, error) {
	ret := _m.Called(activities)

	var r0 *entity.Activities
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Activities) (*entity.Activities, error)); ok {
		return rf(activities)
	}
	if rf, ok := ret.Get(0).(func(entity.Activities) *entity.Activities); ok {
		r0 = rf(activities)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Activities)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Activities) error); ok {
		r1 = rf(activities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryActivities interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryActivities creates a new instance of IRepositoryActivities. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryActivities(t mockConstructorTestingTNewIRepositoryActivities) *IRepositoryActivities {
	mock := &IRepositoryActivities{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}