package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminUsecase "yego/internal/usecases/admin"
)

// NewListOrdersHandler creates a handler for listing all orders
func NewListOrdersHandler(usecase adminUsecase.ListOrdersUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, appErr := usecase.Execute(c)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
