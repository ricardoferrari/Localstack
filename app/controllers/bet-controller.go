package betControllers

import (
	"log"

	"github.com/gin-gonic/gin"
	betUseCase "github.com/ricardoferrari/localstack/usecases"
	"go.uber.org/fx"
)

func NewBetController(lc fx.Lifecycle, logger *log.Logger, httpServer *gin.Engine, useCase betUseCase.MessageUseCaseInterface) error {
	api := httpServer.Group("/api")
	api.GET("/message", func(c *gin.Context) {
		c.JSON(200, useCase.GetMessages())
	})

	// httpServer.Run()
	return nil
}
