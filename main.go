package main

import (
	"log"
	"os"

	"github.com/dealense7/product-app/app/handlers"
	"github.com/dealense7/product-app/app/repositories"
	"github.com/dealense7/product-app/app/services"
	"github.com/dealense7/product-app/utils"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() (*gin.Engine, error) {
		return gin.Default(), nil
	})

	container.Provide(utils.BuildDSN)
	container.Provide(utils.NewDB)

	// Product dependencies
	container.Provide(repositories.NewMySQLProductRepository)
	container.Provide(services.NewProductService)

	// Currency dependencies
	container.Provide(repositories.NewMySQLCurrencyRepository)
	container.Provide(services.NewCurrencyService)

	// Gas dependencies
	container.Provide(repositories.NewMySQLGasRepository)
	container.Provide(services.NewGasService)

	container.Provide(handlers.NewWebHandler)

	container.Provide(zap.NewProduction)

	return container
}

func main() {
	container := BuildContainer()

	// Invoke the container to resolve dependencies and start the server.
	err := container.Invoke(func(
		engine *gin.Engine,
		handler *handlers.WebHandler,
		db *sqlx.DB,
		logger *zap.Logger,
	) {
		// Serve static assets
		engine.Static("/static", "./static")

		// Load HTML templates from specified directory
		engine.LoadHTMLGlob("resources/templates/**/*.html")

		// Define application routes
		engine.GET("/", handler.GetProducts)

		// Determine port from environment or use default 8080
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		// Log the server start and run it
		logger.Info("Starting server", zap.String("port", port))
		if err := engine.Run(":" + port); err != nil {
			logger.Fatal("Server encountered an error", zap.Error(err))
		}
	})
	// If the container fails to resolve dependencies, log the detailed error.
	if err != nil {
		log.Fatalf("Failed to invoke container: %+v", err)
	}
}
