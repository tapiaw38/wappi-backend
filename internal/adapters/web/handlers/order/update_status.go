package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	orderUsecase "wappi/internal/usecases/order"
)

type UpdateStatusInput struct {
	Status string `json:"status" binding:"required"`
}

// NewUpdateStatusHandler creates a handler for updating order status
func NewUpdateStatusHandler(usecase orderUsecase.UpdateStatusUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var input UpdateStatusInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, id, orderUsecase.UpdateStatusInput{
			Status: input.Status,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
