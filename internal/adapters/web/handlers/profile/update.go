package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	profileUsecase "wappi/internal/usecases/profile"
)

type UpdateProfileInput struct {
	PhoneNumber string  `json:"phone_number"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

// NewUpdateHandler creates a handler for updating a profile
func NewUpdateHandler(usecase profileUsecase.UpdateProfileUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var input UpdateProfileInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, id, profileUsecase.UpdateProfileInput{
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

		c.JSON(http.StatusOK, output)
	}
}
