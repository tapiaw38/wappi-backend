package order

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	orderUsecase "yego/internal/usecases/order"
)

type CreateInput struct {
	ProfileID    string `json:"profile_id" binding:"required"`
	ETA          string `json:"eta"`
	SecurityCode string `json:"security_code"`
}

// NewCreateHandler creates a handler for creating orders
func NewCreateHandler(usecase orderUsecase.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		var token string
		if authHeader != "" {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		output, appErr := usecase.Execute(c, orderUsecase.CreateInput{
			ProfileID:    input.ProfileID,
			ETA:          input.ETA,
			SecurityCode: input.SecurityCode,
			Token:        token,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
