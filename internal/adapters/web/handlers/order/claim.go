package order

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"wappi/internal/adapters/web/middlewares"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	orderUsecase "wappi/internal/usecases/order"
)

type ClaimRequestBody struct {
	SecurityCode string `json:"security_code"`
}

// NewClaimHandler creates a handler for claiming orders via token
// This endpoint requires authentication - user_id comes from JWT context
func NewClaimHandler(usecase orderUsecase.ClaimUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		if token == "" {
			appErr := apperrors.NewApplicationError(mappings.OrderTokenNotFoundError, nil)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		// Get user_id from JWT context (set by auth middleware)
		userID, exists := middlewares.GetUserIDFromContext(c)
		if !exists {
			appErr := apperrors.NewApplicationError(mappings.UnauthorizedError, nil)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		// Parse request body for security_code
		var body ClaimRequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			// If no body provided, continue without security_code (backward compatibility)
			body.SecurityCode = ""
		}

		// Extract auth token from header
		authHeader := c.GetHeader("Authorization")
		var authToken string
		if authHeader != "" {
			authToken = strings.TrimPrefix(authHeader, "Bearer ")
		}

		output, appErr := usecase.Execute(c, orderUsecase.ClaimInput{
			Token:        token,
			UserID:       userID,
			SecurityCode: body.SecurityCode,
			AuthToken:    authToken,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
