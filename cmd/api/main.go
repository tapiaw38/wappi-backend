package main

import (
	"log"

	"wappi/internal/adapters/datasources"
	"wappi/internal/adapters/web"
	"wappi/internal/adapters/web/middlewares"
	"wappi/internal/platform/appcontext"
	"wappi/internal/platform/config"
	"wappi/internal/platform/database"
	"wappi/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.GetInstance()
	log.Printf("Starting Order Tracking API on port %s", cfg.ServerPort)

	// Initialize database
	db := database.GetInstance()

	// Run migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create datasources
	ds := datasources.CreateDatasources(db)

	// Create context factory for dependency injection
	contextFactory := appcontext.NewFactory(ds, cfg)

	// Initialize use cases
	useCases := usecases.CreateUsecases(contextFactory)

	// Setup Gin
	gin.SetMode(cfg.GinMode)
	app := gin.Default()

	// Apply CORS middleware
	app.Use(middlewares.CORSMiddleware())

	// Register routes
	web.RegisterRoutes(app, useCases, cfg.FrontendURL)

	// Health check endpoint
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	if err := app.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
