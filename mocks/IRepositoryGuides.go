// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-user/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryGuides is an autogenerated mock type for the IRepositoryGuides type
type IRepositoryGuides struct {
	mock.Mock
}

// DeleteGuides provides a mock function with given fields: idGuides
func (_m *IRepositoryGuides) DeleteGuides(idGuides int64) error {
	ret := _m.Called(idGuides)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(idGuides)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindGuides provides a mock function with given fields: filter
func (_m *IRepositoryGuides) FindGuides(filter map[string]interface{}) (*[]entity.Guides, error) {
	ret := _m.Called(filter)

	var r0 *[]entity.Guides
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) (*[]entity.Guides, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *[]entity.Guides); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Guides)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindGuidesOne provides a mock function with given fields: idGuide
func (_m *IRepositoryGuides) FindGuidesOne(idGuide int64) (*entity.Guides, error) {
	ret := _m.Called(idGuide)

	var r0 *entity.Guides
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*entity.Guides, error)); ok {
		return rf(idGuide)
	}
	if rf, ok := ret.Get(0).(func(int64) *entity.Guides); ok {
		r0 = rf(idGuide)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Guides)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idGuide)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertGuides provides a mock function with given fields: guides
func (_m *IRepositoryGuides) InsertGuides(guides entity.Guides) (*entity.Guides, error) {
	ret := _m.Called(guides)

	var r0 *entity.Guides
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Guides) (*entity.Guides, error)); ok {
		return rf(guides)
	}
	if rf, ok := ret.Get(0).(func(entity.Guides) *entity.Guides); ok {
		r0 = rf(guides)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Guides)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Guides) error); ok {
		r1 = rf(guides)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryGuides interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryGuides creates a new instance of IRepositoryGuides. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryGuides(t mockConstructorTestingTNewIRepositoryGuides) *IRepositoryGuides {
	mock := &IRepositoryGuides{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
