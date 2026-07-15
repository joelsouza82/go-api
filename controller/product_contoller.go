package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(useCase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: useCase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "ProductId is required",
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": response.Message})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ProductId must be a valid integer",
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": response.Message})
		return
	}

	products, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == nil && products == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": response.Message})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
