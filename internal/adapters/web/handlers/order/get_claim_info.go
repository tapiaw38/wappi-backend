package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "wappi/internal/platform/errors"
	"wappi/internal/platform/errors/mappings"
	orderUsecase "wappi/internal/usecases/order"
)

// NewGetClaimInfoHandler creates a handler for getting order info from claim token
func NewGetClaimInfoHandler(usecase orderUsecase.GetClaimInfoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		if token == "" {
			appErr := apperrors.NewApplicationError(mappings.OrderTokenNotFoundError, nil)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, orderUsecase.GetClaimInfoInput{
			Token: token,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
