package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminUsecase "yego/internal/usecases/admin"
)

// NewDeleteImportHandler creates a handler for deleting an import record
func NewDeleteImportHandler(usecase adminUsecase.DeleteImportUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if appErr := usecase.Execute(c, id); appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
