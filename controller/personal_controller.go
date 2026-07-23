package controller

import (
	"go-api/usecase"
	"net/http"

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
