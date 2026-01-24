package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wappi/internal/adapters/web/middlewares"
	orderUsecase "wappi/internal/usecases/order"
)

// NewListMyHandler creates a handler for listing current user's orders
func NewListMyHandler(usecase orderUsecase.ListMyOrdersUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := middlewares.GetUserIDFromContext(c)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
			return
		}

		output, appErr := usecase.Execute(c.Request.Context(), userID)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
