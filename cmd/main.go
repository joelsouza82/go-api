package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic("Erro de conexão com o banco de dados: " + err.Error())
	}

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	PersonalRepository := repository.NewPersonalRepository(dbConnection)
	PersonalUseCase := usecase.NewPersonalUseCase(PersonalRepository)
	PersonalController := controller.NewPersonalController(PersonalUseCase)

	server.GET("/personals", PersonalController.GetPersonals)
	server.GET("/personal/:personalId", PersonalController.GetPersonalByID)
	server.POST("/personal", PersonalController.CreatePersonal)
	server.PUT("/personal/:personalId", PersonalController.UpdatePersonal)
	server.DELETE("/personal/:personalId", PersonalController.DeletePersonal)

	server.Run(":8000")
}
