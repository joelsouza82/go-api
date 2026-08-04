package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

type LoginRepositoryMock struct {
	mock.Mock
}

func (m *LoginRepositoryMock) GetLogins() ([]model.Login, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return []model.Login{}, args.Error(1)
	}
	return args.Get(0).([]model.Login), nil
}

func (m *LoginRepositoryMock) GetLoginByID(id int) (model.Login, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return model.Login{}, args.Error(1)
	}
	return args.Get(0).(model.Login), nil
}

func (m *LoginRepositoryMock) CreateLogin(login model.Login) (int, error) {
	args := m.Called(login)
	if args.Get(1) != nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int), nil
}

func (m *LoginRepositoryMock) UpdateLogin(login model.Login) (model.Login, error) {
	args := m.Called(login)
	if args.Get(1) != nil {
		return model.Login{}, args.Error(1)
	}
	return args.Get(0).(model.Login), nil
}

func (m *LoginRepositoryMock) DeleteLogin(id int) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
