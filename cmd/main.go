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

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	PersonalRepository := repository.NewPersonalRepository(dbConnection)
	PersonalUseCase := usecase.NewPersonalUseCase(PersonalRepository) // Nenhuma mudança necessária aqui, Go lida com isso.
	PersonalController := controller.NewPersonalController(PersonalUseCase)

	server.GET("/personals", PersonalController.GetPersonals)
	server.POST("/personal", PersonalController.CreatePersonal)
	server.PUT("/personal/:personalId", PersonalController.UpdatePersonal)
	server.DELETE("/personal/:personalId", PersonalController.DeletePersonal)

	server.Run(":8000")
}
