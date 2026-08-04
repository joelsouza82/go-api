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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginController_GetLogins(t *testing.T) {
	mockUseCase := new(mocks.LoginUseCaseMock)
	controller := NewLoginController(mockUseCase)

	router := setupRouter()
	router.GET("/login", controller.GetLogins)

	loginMock1 := model.Login{
		ID:       1,
		Email:    "user1@example.com",
		Password: "pass1",
	}

	loginMock2 := model.Login{
		ID:       2,
		Email:    "user2@example.com",
		Password: "pass2",
	}

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetLogins").Return([]model.Login{loginMock1, loginMock2}, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseLogins []model.Login
		err := json.Unmarshal(w.Body.Bytes(), &responseLogins)
		assert.NoError(t, err)
		assert.Len(t, responseLogins, 2)
		assert.Equal(t, loginMock1.ID, responseLogins[0].ID)
		assert.Equal(t, loginMock2.ID, responseLogins[1].ID)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Empty List", func(t *testing.T) {
		mockUseCase.On("GetLogins").Return([]model.Login{}, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseLogins []model.Login
		err := json.Unmarshal(w.Body.Bytes(), &responseLogins)
		assert.NoError(t, err)
		assert.Len(t, responseLogins, 0)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("database error")
		mockUseCase.On("GetLogins").Return([]model.Login{}, internalError).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestLoginController_GetLoginByID(t *testing.T) {
	mockUseCase := new(mocks.LoginUseCaseMock)
	controller := NewLoginController(mockUseCase)

	router := setupRouter()
	router.GET("/login/:loginId", controller.GetLoginByID)

	loginMock := model.Login{
		ID:       1,
		Email:    "test@example.com",
		Password: "pass123",
	}

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetLoginByID", 1).Return(loginMock, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseLogin model.Login
		err := json.Unmarshal(w.Body.Bytes(), &responseLogin)
		assert.NoError(t, err)
		assert.Equal(t, loginMock, responseLogin)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/login/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de login inválido")
	})

	t.Run("Not Found", func(t *testing.T) {
		mockUseCase.On("GetLoginByID", 1).Return(model.Login{}, errors.New("login not found")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de login não encontrado")

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("database error")
		mockUseCase.On("GetLoginByID", 1).Return(model.Login{}, internalError).Once()

		req, _ := http.NewRequest(http.MethodGet, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestLoginController_CreateLogin(t *testing.T) {
	mockUseCase := new(mocks.LoginUseCaseMock)
	controller := NewLoginController(mockUseCase)

	router := setupRouter()
	router.POST("/login", controller.CreateLogin)

	loginToCreate := model.Login{
		Email:    "newuser@example.com",
		Password: "newpass123",
	}

	t.Run("Success", func(t *testing.T) {
		createdLogin := model.Login{
			ID:       10,
			Email:    "newuser@example.com",
			Password: "newpass123",
		}
		mockUseCase.On("CreateLogin", mock.AnythingOfType("model.Login")).Return(createdLogin, nil).Once()

		body, _ := json.Marshal(loginToCreate)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response model.Login
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 10, response.ID)
		assert.Equal(t, "newuser@example.com", response.Email)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Bad Request - Invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString("{invalid json}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("database error")
		mockUseCase.On("CreateLogin", mock.AnythingOfType("model.Login")).Return(model.Login{}, internalError).Once()

		body, _ := json.Marshal(loginToCreate)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestLoginController_UpdateLogin(t *testing.T) {
	mockUseCase := new(mocks.LoginUseCaseMock)
	controller := NewLoginController(mockUseCase)

	router := setupRouter()
	router.PUT("/login/:loginId", controller.UpdateLogin)

	loginToUpdate := model.Login{
		ID:       1,
		Email:    "updated@example.com",
		Password: "updatedpass",
	}

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("UpdateLogin", mock.AnythingOfType("model.Login")).Return(loginToUpdate, nil).Once()

		body, _ := json.Marshal(loginToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/login/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseLogin model.Login
		err := json.Unmarshal(w.Body.Bytes(), &responseLogin)
		assert.NoError(t, err)
		assert.Equal(t, loginToUpdate.Email, responseLogin.Email)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/login/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de login inválido")
	})

	t.Run("Bad Request - Invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/login/1", bytes.NewBufferString("{invalid json}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockUseCase.On("UpdateLogin", mock.AnythingOfType("model.Login")).Return(model.Login{}, errors.New("login not found")).Once()

		body, _ := json.Marshal(loginToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/login/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de login não encontrado")

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("database error")
		mockUseCase.On("UpdateLogin", mock.AnythingOfType("model.Login")).Return(model.Login{}, internalError).Once()

		body, _ := json.Marshal(loginToUpdate)
		req, _ := http.NewRequest(http.MethodPut, "/login/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())

		mockUseCase.AssertExpectations(t)
	})
}

func TestLoginController_DeleteLogin(t *testing.T) {
	mockUseCase := new(mocks.LoginUseCaseMock)
	controller := NewLoginController(mockUseCase)

	router := setupRouter()
	router.DELETE("/login/:loginId", controller.DeleteLogin)

	loginID := 1

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("DeleteLogin", loginID).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/login/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "ID de login inválido")
	})

	t.Run("Not Found", func(t *testing.T) {
		mockUseCase.On("DeleteLogin", loginID).Return(errors.New("login not found")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Registro de login não encontrado")
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		internalError := errors.New("database error")
		mockUseCase.On("DeleteLogin", loginID).Return(internalError).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/login/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), internalError.Error())
		mockUseCase.AssertExpectations(t)
	})
}
