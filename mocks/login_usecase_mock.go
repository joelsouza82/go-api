package mocks

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

type LoginUseCaseMock struct {
	mock.Mock
}

func (m *LoginUseCaseMock) GetLogins() ([]model.Login, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return []model.Login{}, args.Error(1)
	}
	return args.Get(0).([]model.Login), nil
}

func (m *LoginUseCaseMock) GetLoginByID(id int) (model.Login, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return model.Login{}, args.Error(1)
	}
	return args.Get(0).(model.Login), nil
}

func (m *LoginUseCaseMock) CreateLogin(login model.Login) (model.Login, error) {
	args := m.Called(login)
	if args.Get(1) != nil {
		return model.Login{}, args.Error(1)
	}
	return args.Get(0).(model.Login), nil
}

func (m *LoginUseCaseMock) UpdateLogin(login model.Login) (model.Login, error) {
	args := m.Called(login)
	if args.Get(1) != nil {
		return model.Login{}, args.Error(1)
	}
	return args.Get(0).(model.Login), nil
}

func (m *LoginUseCaseMock) DeleteLogin(id int) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
