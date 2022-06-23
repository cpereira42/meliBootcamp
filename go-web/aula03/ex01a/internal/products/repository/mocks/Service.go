// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() ([]products.Product, error) {
	ret := _m.Called()

	var r0 []products.Product
	if rf, ok := ret.Get(0).(func() []products.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Product)
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

// Store provides a mock function with given fields: name, tipo, count, price
func (_m *Service) Store(name string, tipo string, count int, price float64) (products.Product, error) {
	ret := _m.Called(name, tipo, count, price)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(string, string, int, float64) products.Product); ok {
		r0 = rf(name, tipo, count, price)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, float64) error); ok {
		r1 = rf(name, tipo, count, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, name, tipo, count, price
func (_m *Service) Update(id int, name string, tipo string, count int, price float64) (products.Product, error) {
	ret := _m.Called(id, name, tipo, count, price)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(int, string, string, int, float64) products.Product); ok {
		r0 = rf(id, name, tipo, count, price)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, int, float64) error); ok {
		r1 = rf(id, name, tipo, count, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateName provides a mock function with given fields: id, name
func (_m *Service) UpdateName(id int, name string) (products.Product, error) {
	ret := _m.Called(id, name)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(int, string) products.Product); ok {
		r0 = rf(id, name)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t NewServiceT) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}