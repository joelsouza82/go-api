package usecase

import (
	"errors"
	"go-api/mocks"
	"go-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_GetLoginByID(t *testing.T) {
	mockRepo := new(mocks.LoginRepositoryMock)
	useCase := NewLoginUseCase(mockRepo)

	loginMock := model.Login{
		ID:       1,
		Email:    "test@test.com",
		Password: "pass123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetLoginByID", 1).Return(loginMock, nil).Once()

		login, err := useCase.GetLoginByID(1)

		assert.NoError(t, err)
		assert.Equal(t, loginMock, login)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("login not found")
		mockRepo.On("GetLoginByID", 1).Return(model.Login{}, repoError).Once()

		_, err := useCase.GetLoginByID(1)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestLoginUseCase_GetLogins(t *testing.T) {
	mockRepo := new(mocks.LoginRepositoryMock)
	useCase := NewLoginUseCase(mockRepo)

	loginMock1 := model.Login{
		ID:       1,
		Email:    "user1@test.com",
		Password: "pass1",
	}

	loginMock2 := model.Login{
		ID:       2,
		Email:    "user2@test.com",
		Password: "pass2",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetLogins").Return([]model.Login{loginMock1, loginMock2}, nil).Once()

		logins, err := useCase.GetLogins()

		assert.NoError(t, err)
		assert.Len(t, logins, 2)
		assert.Equal(t, loginMock1, logins[0])
		assert.Equal(t, loginMock2, logins[1])
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("database error")
		mockRepo.On("GetLogins").Return([]model.Login{}, repoError).Once()

		_, err := useCase.GetLogins()

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty List", func(t *testing.T) {
		mockRepo.On("GetLogins").Return([]model.Login{}, nil).Once()

		logins, err := useCase.GetLogins()

		assert.NoError(t, err)
		assert.Len(t, logins, 0)
		mockRepo.AssertExpectations(t)
	})
}

func TestLoginUseCase_CreateLogin(t *testing.T) {
	mockRepo := new(mocks.LoginRepositoryMock)
	useCase := NewLoginUseCase(mockRepo)

	loginToCreate := model.Login{
		Email:    "newuser@test.com",
		Password: "newpass",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("CreateLogin", loginToCreate).Return(10, nil).Once()

		createdLogin, err := useCase.CreateLogin(loginToCreate)

		assert.NoError(t, err)
		assert.Equal(t, 10, createdLogin.ID)
		assert.Equal(t, loginToCreate.Email, createdLogin.Email)
		assert.Equal(t, loginToCreate.Password, createdLogin.Password)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("database error")
		mockRepo.On("CreateLogin", loginToCreate).Return(0, repoError).Once()

		_, err := useCase.CreateLogin(loginToCreate)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestLoginUseCase_UpdateLogin(t *testing.T) {
	mockRepo := new(mocks.LoginRepositoryMock)
	useCase := NewLoginUseCase(mockRepo)

	loginToUpdate := model.Login{
		ID:       1,
		Email:    "updated@test.com",
		Password: "updatedpass",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("UpdateLogin", loginToUpdate).Return(loginToUpdate, nil).Once()

		updatedLogin, err := useCase.UpdateLogin(loginToUpdate)

		assert.NoError(t, err)
		assert.Equal(t, loginToUpdate.ID, updatedLogin.ID)
		assert.Equal(t, loginToUpdate.Email, updatedLogin.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("login not found")
		mockRepo.On("UpdateLogin", loginToUpdate).Return(model.Login{}, repoError).Once()

		_, err := useCase.UpdateLogin(loginToUpdate)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestLoginUseCase_DeleteLogin(t *testing.T) {
	mockRepo := new(mocks.LoginRepositoryMock)
	useCase := NewLoginUseCase(mockRepo)
	loginID := 1

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeleteLogin", loginID).Return(nil).Once()

		err := useCase.DeleteLogin(loginID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("login not found")
		mockRepo.On("DeleteLogin", loginID).Return(repoError).Once()

		err := useCase.DeleteLogin(loginID)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})
}
