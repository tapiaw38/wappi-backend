package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	orderUsecase "wappi/internal/usecases/order"
)

type CreateInput struct {
	ProfileID string `json:"profile_id" binding:"required"`
	ETA       string `json:"eta"`
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

		output, appErr := usecase.Execute(c, orderUsecase.CreateInput{
			ProfileID: input.ProfileID,
			ETA:       input.ETA,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
