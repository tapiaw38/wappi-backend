package web

import (
	"github.com/gin-gonic/gin"
	adminHandler "wappi/internal/adapters/web/handlers/admin"
	orderHandler "wappi/internal/adapters/web/handlers/order"
	profileHandler "wappi/internal/adapters/web/handlers/profile"
	settingsHandler "wappi/internal/adapters/web/handlers/settings"
	"wappi/internal/adapters/web/middlewares"
	"wappi/internal/usecases"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(app *gin.Engine, useCases *usecases.Usecases, frontendURL string) {
	api := app.Group("/api")

	// Public order routes (tracking by UUID - no auth needed)
	orders := api.Group("/orders")
	{
		orders.GET("/:id", orderHandler.NewGetHandler(useCases.Order.GetUsecase))
		// Create order with link (admin functionality, no auth for now)
		orders.POST("/create-with-link", orderHandler.NewCreateWithLinkHandler(useCases.Order.CreateWithLinkUsecase, frontendURL))
	}

	// Protected order routes (require auth)
	ordersAuth := api.Group("/orders")
	ordersAuth.Use(middlewares.AuthMiddleware())
	{
		ordersAuth.POST("", orderHandler.NewCreateHandler(useCases.Order.CreateUsecase))
		ordersAuth.PATCH("/:id/status", orderHandler.NewUpdateStatusHandler(useCases.Order.UpdateStatusUsecase))
		// Claim order with token
		ordersAuth.POST("/claim/:token", orderHandler.NewClaimHandler(useCases.Order.ClaimUsecase))
		// List current user's orders
		ordersAuth.GET("/my", orderHandler.NewListMyHandler(useCases.Order.ListMyOrdersUsecase))
	}

	// Public profile routes (token-based access)
	profiles := api.Group("/profiles")
	{
		profiles.GET("/validate/:token", profileHandler.NewValidateTokenHandler(useCases.Profile.ValidateTokenUsecase))
		profiles.POST("/complete", profileHandler.NewCompleteProfileHandler(useCases.Profile.CompleteProfileUsecase))
		profiles.GET("/:id", profileHandler.NewGetHandler(useCases.Profile.GetUsecase))
		profiles.PUT("/:id", profileHandler.NewUpdateHandler(useCases.Profile.UpdateUsecase))
	}

	// Protected profile routes (require auth)
	profilesAuth := api.Group("/profiles")
	profilesAuth.Use(middlewares.AuthMiddleware())
	{
		profilesAuth.POST("/generate-link", profileHandler.NewGenerateLinkHandler(useCases.Profile.GenerateLinkUsecase))
		profilesAuth.GET("/check-completed", profileHandler.NewCheckCompletedHandler(useCases.Profile.CheckCompletedUsecase))
	}

	// Settings routes (public for reading, could add auth for updating)
	settings := api.Group("/settings")
	{
		settings.GET("", settingsHandler.NewGetHandler(useCases.Settings.GetUsecase))
		settings.PUT("", settingsHandler.NewUpdateHandler(useCases.Settings.UpdateUsecase))
		settings.POST("/calculate-delivery", settingsHandler.NewCalculateDeliveryFeeHandler(useCases.Settings.CalculateDeliveryFeeUsecase))
	}

	// Admin routes (require auth)
	admin := api.Group("/admin")
	admin.Use(middlewares.AuthMiddleware())
	{
		admin.GET("/profiles", adminHandler.NewListProfilesHandler(useCases.Admin.ListProfilesUsecase))
		admin.GET("/orders", adminHandler.NewListOrdersHandler(useCases.Admin.ListOrdersUsecase))
		admin.PUT("/orders/:id", adminHandler.NewUpdateOrderHandler(useCases.Admin.UpdateOrderUsecase))
	}
}
