package main

import (
	"log"

	"wappi/internal/adapters/datasources"
	"wappi/internal/adapters/web"
	websocketHandler "wappi/internal/adapters/web/handlers/websocket"
	"wappi/internal/adapters/web/integrations"
	"wappi/internal/adapters/web/middlewares"
	"wappi/internal/platform/appcontext"
	"wappi/internal/platform/config"
	"wappi/internal/platform/database"
	"wappi/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetInstance()
	log.Printf("Starting Order Tracking API on port %s", cfg.ServerPort)

	db := database.GetInstance()

	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	ds := datasources.CreateDatasources(db)
	integrations := integrations.CreateIntegration(cfg)
	contextFactory := appcontext.NewFactory(ds, integrations, cfg)

	useCases := usecases.CreateUsecases(contextFactory)

	gin.SetMode(cfg.GinMode)
	app := gin.Default()

	app.Use(middlewares.CORSMiddleware())

	hub := integrations.WebSocket.GetHub()
	wsHandler := websocketHandler.NewHandler(hub)
	web.RegisterRoutes(app, useCases, cfg.FrontendURL, wsHandler, contextFactory)

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := app.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
