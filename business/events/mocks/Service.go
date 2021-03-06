// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	events "yukevent/business/events"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AllEvent provides a mock function with given fields:
func (_m *Service) AllEvent() ([]events.Domain, error) {
	ret := _m.Called()

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func() []events.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: orgID, domain
func (_m *Service) Create(orgID int, domain *events.Domain) (events.Domain, error) {
	ret := _m.Called(orgID, domain)

	var r0 events.Domain
	if rf, ok := ret.Get(0).(func(int, *events.Domain) events.Domain); ok {
		r0 = rf(orgID, domain)
	} else {
		r0 = ret.Get(0).(events.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *events.Domain) error); ok {
		r1 = rf(orgID, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: orgID, id
func (_m *Service) Delete(orgID int, id int) (string, error) {
	ret := _m.Called(orgID, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(orgID, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(orgID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventByID provides a mock function with given fields: id
func (_m *Service) EventByID(id int) (events.Domain, error) {
	ret := _m.Called(id)

	var r0 events.Domain
	if rf, ok := ret.Get(0).(func(int) events.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(events.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventByIdOrganizer provides a mock function with given fields: orgzID
func (_m *Service) EventByIdOrganizer(orgzID int) ([]events.Domain, error) {
	ret := _m.Called(orgzID)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(int) []events.Domain); ok {
		r0 = rf(orgzID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orgzID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyEventByOrganizer provides a mock function with given fields: orgID
func (_m *Service) MyEventByOrganizer(orgID int) ([]events.Domain, error) {
	ret := _m.Called(orgID)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(int) []events.Domain); ok {
		r0 = rf(orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: orgID, evID, domain
func (_m *Service) Update(orgID int, evID int, domain *events.Domain) (events.Domain, error) {
	ret := _m.Called(orgID, evID, domain)

	var r0 events.Domain
	if rf, ok := ret.Get(0).(func(int, int, *events.Domain) events.Domain); ok {
		r0 = rf(orgID, evID, domain)
	} else {
		r0 = ret.Get(0).(events.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, *events.Domain) error); ok {
		r1 = rf(orgID, evID, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
