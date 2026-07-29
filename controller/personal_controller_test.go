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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
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
