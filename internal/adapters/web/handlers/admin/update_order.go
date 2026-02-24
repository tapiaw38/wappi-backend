package admin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	adminUsecase "yego/internal/usecases/admin"
)

type UpdateOrderInput struct {
	Status        *string           `json:"status,omitempty"`
	StatusMessage *string           `json:"status_message,omitempty"`
	ETA           *string           `json:"eta,omitempty"`
	Data          *domain.OrderData `json:"data,omitempty"`
}

// NewUpdateOrderHandler creates a handler for updating an order
func NewUpdateOrderHandler(usecase adminUsecase.UpdateOrderUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var input UpdateOrderInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		authHeader := c.GetHeader("Authorization")
		var token string
		if authHeader != "" {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		output, appErr := usecase.Execute(c, id, adminUsecase.UpdateOrderInput{
			Status:        input.Status,
			StatusMessage: input.StatusMessage,
			ETA:           input.ETA,
			Data:          input.Data,
			Token:         token,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
