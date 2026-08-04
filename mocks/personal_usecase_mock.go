package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

type PersonalUseCaseMock struct {
	mock.Mock
}

func (m *PersonalUseCaseMock) GetPersonalByID(id int) (model.Personal, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), nil
}

func (m *PersonalUseCaseMock) GetPersonals() ([]model.Personal, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return []model.Personal{}, args.Error(1)
	}
	return args.Get(0).([]model.Personal), nil
}

func (m *PersonalUseCaseMock) CreatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	if args.Get(1) != nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), nil
}

func (m *PersonalUseCaseMock) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	if args.Get(1) != nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), nil
}

func (m *PersonalUseCaseMock) DeletePersonal(id int) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
