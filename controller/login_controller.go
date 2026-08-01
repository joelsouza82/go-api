package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	loginUsecase usecase.LoginUseCase
}

func NewLoginController(useCase usecase.LoginUseCase) loginController {
	return loginController{
		loginUsecase: useCase,
	}
}

func (l *loginController) GetLogins(ctx *gin.Context) {
	logins, err := l.loginUsecase.GetLogins()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, logins)
}

func (l *loginController) GetLoginByID(ctx *gin.Context) {
	idStr := ctx.Param("loginId")
	loginId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de login inválido"})
		return
	}

	login, err := l.loginUsecase.GetLoginByID(loginId)
	if err != nil {
		if err.Error() == "login not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Registro de login não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, login)
}

func (l *loginController) CreateLogin(ctx *gin.Context) {
	var login model.Login
	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedLogin, err := l.loginUsecase.CreateLogin(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedLogin)
}

func (l *loginController) UpdateLogin(ctx *gin.Context) {
	idStr := ctx.Param("loginId")
	loginId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de login inválido"})
		return
	}

	var login model.Login
	if err := ctx.BindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login.ID = loginId

	updatedLogin, err := l.loginUsecase.UpdateLogin(login)
	if err != nil {
		if err.Error() == "login not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Registro de login não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedLogin)
}

func (l *loginController) DeleteLogin(ctx *gin.Context) {
	idStr := ctx.Param("loginId")
	loginId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de login inválido"})
		return
	}

	if err := l.loginUsecase.DeleteLogin(loginId); err != nil {
		if err.Error() == "login not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Registro de login não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
