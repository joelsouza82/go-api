package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

type PersonalRepositoryMock struct {
	mock.Mock
}

func (m *PersonalRepositoryMock) GetPersonalByID(id int) (model.Personal, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), nil
}

func (m *PersonalRepositoryMock) GetPersonals() ([]model.Personal, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return []model.Personal{}, args.Error(1)
	}
	return args.Get(0).([]model.Personal), nil
}

func (m *PersonalRepositoryMock) CreatePersonal(personal model.Personal) (int, error) {
	args := m.Called(personal)
	if args.Get(1) != nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int), nil
}

func (m *PersonalRepositoryMock) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	if args.Get(1) != nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), nil
}

func (m *PersonalRepositoryMock) DeletePersonal(id int) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
