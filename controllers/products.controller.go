package controllers

import (
	"lumo-pos/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductsController struct {
	DB *gorm.DB
}

func NewProductsController(db *gorm.DB) ProductsController {
	return ProductsController{db}
}

func (pc *ProductsController) CreateProduct(ctx *gin.Context) {
	var payload *models.CreateProductRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newProduct := models.Products{
		Name:      payload.Name,
		Price:     payload.Price,
		Quantity:  payload.Quantity,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newProduct)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Producto ya existe"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProduct})
}

func (pc *ProductsController) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	var payload *models.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updateProduct models.Products
	result := pc.DB.First(&updateProduct, "id = ?", productId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Producto no encontrado"})
		return
	}
	now := time.Now()
	productToUpdate := models.Products{
		Name:      payload.Name,
		Price:     payload.Price,
		Quantity:  payload.Quantity,
		CreatedAt: updateProduct.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updateProduct).Updates(productToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": productToUpdate})
}

func (pc *ProductsController) FindProductById(ctx *gin.Context) {
	productId := ctx.Param("id")

	var product models.Products
	result := pc.DB.First(&product, "id = ?", productId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Producto no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": product})
}

func (pc *ProductsController) FindAllProducts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var products []models.Products
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&products)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": results.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(products), "data": products})
}

func (pc *ProductsController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	result := pc.DB.Delete(&models.Products{}, "id = ?", productId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Producto no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Producto [" + productId + "] eliminado correctamente."})
}
