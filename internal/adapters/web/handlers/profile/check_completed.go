package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yego/internal/adapters/web/middlewares"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	profileUsecase "yego/internal/usecases/profile"
)

// NewCheckCompletedHandler creates a handler for checking if profile is completed
// User ID is extracted from JWT token (set by AuthMiddleware)
func NewCheckCompletedHandler(usecase profileUsecase.CheckCompletedUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from JWT context (set by AuthMiddleware)
		userID, exists := middlewares.GetUserIDFromContext(c)
		if !exists || userID == "" {
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
