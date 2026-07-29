package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

// Mock para PersonalRepository
type PersonalRepositoryMock struct {
	mock.Mock
}

func (m *PersonalRepositoryMock) GetPersonals() ([]model.Personal, error) {
	args := m.Called()
	return args.Get(0).([]model.Personal), args.Error(1)
}

func (m *PersonalRepositoryMock) CreatePersonal(personal model.Personal) (int, error) {
	args := m.Called(personal)
	return args.Int(0), args.Error(1)
}

func (m *PersonalRepositoryMock) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	args := m.Called(personal)
	if args.Get(0) == nil {
		return model.Personal{}, args.Error(1)
	}
	return args.Get(0).(model.Personal), args.Error(1)
}

func (m *PersonalRepositoryMock) DeletePersonal(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
