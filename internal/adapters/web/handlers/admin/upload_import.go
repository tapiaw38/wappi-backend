package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	adminUsecase "yego/internal/usecases/admin"
)

// NewUploadImportHandler creates a handler for uploading an Excel file
func NewUploadImportHandler(usecase adminUsecase.UploadImportUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			appErr := apperrors.NewApplicationError(mappings.ImportFileParseError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			appErr := apperrors.NewApplicationError(mappings.ImportFileParseError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}
		defer file.Close()

		output, appErr := usecase.Execute(c, file)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
