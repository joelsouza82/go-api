package usecase

import (
	"errors"
	"go-api/mocks"
	"go-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonalUseCase_UpdatePersonal(t *testing.T) {
	mockRepo := new(mocks.PersonalRepositoryMock)
	// Agora podemos passar o mock diretamente, pois o UseCase espera a interface.
	useCase := NewPersonalUseCase(mockRepo)

	personalToUpdate := model.Personal{
		ID:    1,
		Email: "updated@test.com",
	}

	t.Run("Success", func(t *testing.T) {
		// Configura o mock para esperar a chamada de UpdatePersonal
		mockRepo.On("UpdatePersonal", personalToUpdate).Return(personalToUpdate, nil).Once()

		updatedPersonal, err := useCase.UpdatePersonal(personalToUpdate)

		assert.NoError(t, err)
		assert.Equal(t, personalToUpdate.ID, updatedPersonal.ID)
		assert.Equal(t, personalToUpdate.Email, updatedPersonal.Email)

		// Verifica se o método esperado foi chamado no mock
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("database connection failed")

		// Configura o mock para retornar um erro
		mockRepo.On("UpdatePersonal", personalToUpdate).Return(model.Personal{}, repoError).Once()

		_, err := useCase.UpdatePersonal(personalToUpdate)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)

		mockRepo.AssertExpectations(t)
	})
}

func TestPersonalUseCase_DeletePersonal(t *testing.T) {
	mockRepo := new(mocks.PersonalRepositoryMock)
	useCase := NewPersonalUseCase(mockRepo)
	personalID := 1

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeletePersonal", personalID).Return(nil).Once()

		err := useCase.DeletePersonal(personalID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repoError := errors.New("record not found")
		mockRepo.On("DeletePersonal", personalID).Return(repoError).Once()

		err := useCase.DeletePersonal(personalID)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		mockRepo.AssertExpectations(t)
	})
}
