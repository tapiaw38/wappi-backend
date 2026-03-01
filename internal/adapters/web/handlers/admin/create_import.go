package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	adminUsecase "yego/internal/usecases/admin"
)

// NewCreateImportHandler creates a handler for creating a single import record
func NewCreateImportHandler(usecase adminUsecase.CreateImportUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input adminUsecase.CreateImportInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, input)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
