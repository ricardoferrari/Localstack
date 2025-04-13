package bootstrap

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, logger *log.Logger) *gin.Engine {
	logger.Print("Executing NewHTTPServer.")
	router := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Print("Starting HTTP server.")
			go func() {
				// service connections
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				logger.Fatal("Server Shutdown:", err)
			}

			logger.Print("Shutting down HTTP server.")
			return nil
		},
	})

	return router
}
