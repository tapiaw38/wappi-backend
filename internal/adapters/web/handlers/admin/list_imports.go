package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminUsecase "yego/internal/usecases/admin"
)

// NewListImportsHandler creates a handler for listing import records
func NewListImportsHandler(usecase adminUsecase.ListImportsUsecase) gin.HandlerFunc {
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
