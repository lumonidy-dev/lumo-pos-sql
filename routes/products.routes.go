package routes

import (
	"lumo-pos/controllers"
	"lumo-pos/middleware"

	"github.com/gin-gonic/gin"
)

type ProductsRouteController struct {
	productsController controllers.ProductsController
}

func NewRouteProductsController(ProductsController controllers.ProductsController) ProductsRouteController {
	return ProductsRouteController{ProductsController}
}

func (pc *ProductsRouteController) ProductsRoute(rg *gin.RouterGroup) {
	router := rg.Group("products")
	router.Use(middleware.DeserializeUser())
	router.GET("/", pc.productsController.FindAllProducts)
	router.GET("/:id", pc.productsController.FindProductById)
	router.POST("/", pc.productsController.CreateProduct)
	router.PUT("/:id", pc.productsController.UpdateProduct)
	router.DELETE("/:id", pc.productsController.DeleteProduct)
}
