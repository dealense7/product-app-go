package main

import (
	"github.com/dealense7/product-app/app/handlers"
	"github.com/dealense7/product-app/app/repositories"
	"github.com/dealense7/product-app/app/services"
	"github.com/dealense7/product-app/utils"
	"github.com/jmoiron/sqlx"
	"os"

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
	container.Provide(repositories.NewMySQLProductRepository)
	container.Provide(services.NewProductService)
	container.Provide(handlers.NewProductHandler)
	container.Provide(zap.NewProduction)

	return container
}
func main() {
	container := BuildContainer()

	err := container.Invoke(func(
		engine *gin.Engine,
		handler *handlers.ProductHandler,
		db *sqlx.DB,
		logger *zap.Logger,
	) {
		// Load Assets
		engine.Static("/static", "./static")

		// Load templates with explicit path
		engine.LoadHTMLGlob("resource/templates/**/*.html")

		// Routes
		engine.GET("/", handler.GetProducts)

		// Server setup
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		logger.Info("Starting server", zap.String("port", port))
		engine.Run(":" + port)
	})

	if err != nil {
		panic(err)
	}
}
