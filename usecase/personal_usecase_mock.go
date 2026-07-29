package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

type PersonalUseCaseMock struct {
	mock.Mock
}

func (m *PersonalUseCaseMock) GetPersonals() ([]model.Personal, error) {
	args := m.Called()
	return args.Get(0).([]model.Personal), args.Error(1)
}

func (m *PersonalUseCaseMock) CreatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	return args.Get(0).(model.Personal), args.Error(1)
}

func (m *PersonalUseCaseMock) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	if args.Get(0) == nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), args.Error(1)
}

func (m *PersonalUseCaseMock) DeletePersonal(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
