package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminUsecase "yego/internal/usecases/admin"
)

// NewClearImportsHandler creates a handler that deletes all import records
func NewClearImportsHandler(usecase adminUsecase.ClearImportsUsecase) gin.HandlerFunc {
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
