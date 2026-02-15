package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wappi/internal/adapters/web/middlewares"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	profileUsecase "wappi/internal/usecases/profile"
)

type UpsertProfileInput struct {
	PhoneNumber string  `json:"phone_number"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
}

func NewUpsertHandler(usecase profileUsecase.UpsertProfileUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := middlewares.GetUserIDFromContext(c)
		if !exists || userID == "" {
			appErr := apperrors.NewApplicationError(mappings.UnauthorizedError, nil)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		var input UpsertProfileInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, userID, profileUsecase.UpsertProfileInput{
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
