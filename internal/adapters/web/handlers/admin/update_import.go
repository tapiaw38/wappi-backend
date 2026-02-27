package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	adminUsecase "yego/internal/usecases/admin"
)

// NewUpdateImportHandler creates a handler for updating an import record
func NewUpdateImportHandler(usecase adminUsecase.UpdateImportUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var input adminUsecase.UpdateImportInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, id, input)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
