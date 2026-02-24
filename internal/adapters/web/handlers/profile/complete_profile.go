package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	profileUsecase "yego/internal/usecases/profile"
)

type CompleteProfileInput struct {
	Token       string  `json:"token" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Address     string  `json:"address"`
}

// NewCompleteProfileHandler creates a handler for completing profiles
func NewCompleteProfileHandler(usecase profileUsecase.CompleteProfileUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CompleteProfileInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, profileUsecase.CompleteProfileInput{
			Token:       input.Token,
			PhoneNumber: input.PhoneNumber,
			Longitude:   input.Longitude,
			Latitude:    input.Latitude,
			Address:     input.Address,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
