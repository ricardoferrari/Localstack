package productControllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	productUseCase "github.com/ricardoferrari/gorest/usecases"
	"go.uber.org/fx"
)

func NewProductController(lc fx.Lifecycle, logger *log.Logger, httpServer *gin.Engine, useCase productUseCase.ProductListInterface) error {
	api := httpServer.Group("/api")
	api.GET("/products", func(c *gin.Context) {
		c.JSON(200, useCase.GetProducts())
	})
	api.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		productID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid product ID"})
			return
		}
		c.JSON(200, useCase.GetProduct(productID))
	})

	// httpServer.Run()
	return nil
}
