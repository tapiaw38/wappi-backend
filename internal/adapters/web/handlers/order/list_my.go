package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yego/internal/adapters/web/middlewares"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	orderUsecase "yego/internal/usecases/order"
)

// NewListMyHandler creates a handler for listing current user's orders
func NewListMyHandler(usecase orderUsecase.ListMyOrdersUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := middlewares.GetUserIDFromContext(c)
		if !exists {
			appErr := apperrors.NewApplicationError(mappings.UnauthorizedError, nil)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, userID)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
