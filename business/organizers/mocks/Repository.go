// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	organizers "yukevent/business/organizers"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, password
func (_m *Repository) Login(email string, password string) (organizers.Domain, error) {
	ret := _m.Called(email, password)

	var r0 organizers.Domain
	if rf, ok := ret.Get(0).(func(string, string) organizers.Domain); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(organizers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: domain
func (_m *Repository) Register(domain *organizers.Domain) (organizers.Domain, error) {
	ret := _m.Called(domain)

	var r0 organizers.Domain
	if rf, ok := ret.Get(0).(func(*organizers.Domain) organizers.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(organizers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*organizers.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
