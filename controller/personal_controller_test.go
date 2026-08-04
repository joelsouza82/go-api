package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-api/mocks"
	"go-api/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestPersonalController_GetPersonals(t *testing.T) {
	mockUseCase := new(mocks.PersonalUseCaseMock)
	controller := NewPersonalController(mockUseCase)

	router := setupRouter()
	router.GET("/personal", controller.GetPersonals)

	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	personalMock1 := model.Personal{
		ID:        1,
		Name:      "User One",
		Email:     "user1@example.com",
		BirthDate: birthDate,
		LoginId:   1,
	}

	personalMock2 := model.Personal{
		ID:        2,
		Name:      "User Two",
		Email:     "user2@example.com",
		BirthDate: birthDate,
		LoginId:   2,
	}

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetPersonals").Return([]model.Personal{personalMock1, personalMock2}, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responsePersonals []model.Personal
		err := json.Unmarshal(w.Body.Bytes(), &responsePersonals)
		assert.NoError(t, err)
		assert.Len(t, responsePersonals, 2)
		assert.Equal(t, personalMock1.ID, responsePersonals[0].ID)
		assert.Equal(t, personalMock2.ID, responsePersonals[1].ID)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Empty List", func(t *testing.T) {
		mockUseCase.On("GetPersonals").Return([]model.Personal{}, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responsePersonals []model.Personal
		err := json.Unmarshal(w.Body.Bytes(), &responsePersonals)
		assert.NoError(t, err)
		assert.Len(t, responsePersonals, 0)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("some internal error")
		mockUseCase.On("GetPersonals").Return([]model.Personal{}, internalError).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestPersonalController_GetPersonalByID(t *testing.T) {
	mockUseCase := new(mocks.PersonalUseCaseMock)
	controller := NewPersonalController(mockUseCase)

	router := setupRouter()
	router.GET("/personal/:personalId", controller.GetPersonalByID)

	personalMock := model.Personal{
		ID:    1,
		Email: "test.user@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetPersonalByID", 1).Return(personalMock, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responsePersonal model.Personal
		err := json.Unmarshal(w.Body.Bytes(), &responsePersonal)
		assert.NoError(t, err)
		assert.Equal(t, personalMock, responsePersonal)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/personal/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de personal inválido")
	})

	t.Run("Not Found", func(t *testing.T) {
		mockUseCase.On("GetPersonalByID", 1).Return(model.Personal{}, errors.New("personal not found")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de personal não encontrado")

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("some internal error")
		mockUseCase.On("GetPersonalByID", 1).Return(model.Personal{}, internalError).Once()

		req, _ := http.NewRequest(http.MethodGet, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestPersonalController_UpdatePersonal(t *testing.T) {
	mockUseCase := new(mocks.PersonalUseCaseMock)
	controller := NewPersonalController(mockUseCase)

	router := setupRouter()
	router.PUT("/personal/:personalId", controller.UpdatePersonal)

	personalToUpdate := model.Personal{
		ID:    1,
		Email: "updated.user@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		// Setup mock
		mockUseCase.On("UpdatePersonal", mock.AnythingOfType("model.Personal")).Return(personalToUpdate, nil).Once()

		// Cria o corpo da requisição
		body, _ := json.Marshal(personalToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/personal/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		// Grava a resposta
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Asserts
		assert.Equal(t, http.StatusOK, w.Code)

		var responsePersonal model.Personal
		err := json.Unmarshal(w.Body.Bytes(), &responsePersonal)
		assert.NoError(t, err)
		assert.Equal(t, personalToUpdate.Email, responsePersonal.Email)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/personal/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de personal inválido")
	})

	t.Run("Bad Request - Invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/personal/1", bytes.NewBufferString("{invalid json}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Not Found", func(t *testing.T) {
		// Setup mock
		mockUseCase.On("UpdatePersonal", mock.AnythingOfType("model.Personal")).Return(model.Personal{}, errors.New("personal not found")).Once()

		body, _ := json.Marshal(personalToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/personal/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de personal não encontrado")

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		// Setup mock
		internalError := errors.New("some internal error")
		mockUseCase.On("UpdatePersonal", mock.AnythingOfType("model.Personal")).Return(model.Personal{}, internalError).Once()

		body, _ := json.Marshal(personalToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/personal/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestPersonalController_CreatePersonal(t *testing.T) {
	mockUseCase := new(mocks.PersonalUseCaseMock)
	controller := NewPersonalController(mockUseCase)

	router := setupRouter()
	router.POST("/personal", controller.CreatePersonal)

	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	personalToCreate := model.Personal{
		Name:      "New User",
		Email:     "newuser@example.com",
		BirthDate: birthDate,
		LoginId:   1,
	}

	t.Run("Success", func(t *testing.T) {
		createdPersonal := model.Personal{
			ID:        10,
			Name:      "New User",
			Email:     "newuser@example.com",
			BirthDate: birthDate,
			LoginId:   1,
		}
		mockUseCase.On("CreatePersonal", mock.AnythingOfType("model.Personal")).Return(createdPersonal, nil).Once()

		body, _ := json.Marshal(personalToCreate)
		req, _ := http.NewRequest(http.MethodPost, "/personal", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response model.Personal
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 10, response.ID)
		assert.Equal(t, "New User", response.Name)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Bad Request - Invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/personal", bytes.NewBufferString("{invalid json}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("some internal error")
		mockUseCase.On("CreatePersonal", mock.AnythingOfType("model.Personal")).Return(0, internalError).Once()

		body, _ := json.Marshal(personalToCreate)
		req, _ := http.NewRequest(http.MethodPost, "/personal", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestPersonalController_DeletePersonal(t *testing.T) {
	mockUseCase := new(mocks.PersonalUseCaseMock)
	controller := NewPersonalController(mockUseCase)

	router := setupRouter()
	router.DELETE("/personal/:personalId", controller.DeletePersonal)

	personalID := 1

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("DeletePersonal", personalID).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/personal/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de personal inválido")
	})

	t.Run("Not Found", func(t *testing.T) {
		mockUseCase.On("DeletePersonal", personalID).Return(errors.New("personal not found")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de personal não encontrado")
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("some internal error")
		mockUseCase.On("DeletePersonal", personalID).Return(internalError).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/personal/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())
		mockUseCase.AssertExpectations(t)
	})
}
