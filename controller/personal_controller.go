package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type personalController struct {
	personalUsecase usecase.PersonalUseCase
}

func NewPersonalController(useCase usecase.PersonalUseCase) personalController {
	return personalController{
		personalUsecase: useCase,
	}
}

func (p *personalController) GetPersonals(ctx *gin.Context) {
	personals, err := p.personalUsecase.GetPersonals()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, personals)
}

func (p *personalController) CreatePersonal(ctx *gin.Context) {
	var personal model.Personal
	err := ctx.BindJSON(&personal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedPersonal, err := p.personalUsecase.CreatePersonal(personal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedPersonal)
}

func (p *personalController) UpdatePersonal(ctx *gin.Context) {
	idStr := ctx.Param("personalId")
	personalId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de personal inválido"})
		return
	}

	var personal model.Personal
	if err := ctx.BindJSON(&personal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personal.ID = personalId

	updatedPersonal, err := p.personalUsecase.UpdatePersonal(personal)
	if err != nil {
		if err.Error() == "personal not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Registro de personal não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedPersonal)
}

func (p *personalController) DeletePersonal(ctx *gin.Context) {
	idStr := ctx.Param("personalId")
	personalId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de personal inválido"})
		return
	}

	if err := p.personalUsecase.DeletePersonal(personalId); err != nil {
		if err.Error() == "personal not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Registro de personal não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
