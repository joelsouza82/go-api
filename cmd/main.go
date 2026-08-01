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

	LoginRepository := repository.NewLoginRepository(dbConnection)
	LoginUseCase := usecase.NewLoginUseCase(LoginRepository)
	LoginController := controller.NewLoginController(LoginUseCase)

	server.GET("/personals", PersonalController.GetPersonals)
	server.GET("/personal/:personalId", PersonalController.GetPersonalByID)
	server.POST("/personal", PersonalController.CreatePersonal)
	server.PUT("/personal/:personalId", PersonalController.UpdatePersonal)
	server.DELETE("/personal/:personalId", PersonalController.DeletePersonal)

	server.GET("/logins", LoginController.GetLogins)
	server.GET("/login/:loginId", LoginController.GetLoginByID)
	server.POST("/login", LoginController.CreateLogin)
	server.PUT("/login/:loginId", LoginController.UpdateLogin)
	server.DELETE("/login/:loginId", LoginController.DeleteLogin)

	server.Run(":8000")
}
