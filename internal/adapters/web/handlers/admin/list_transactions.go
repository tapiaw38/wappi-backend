package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	adminUsecase "wappi/internal/usecases/admin"
)

// NewListTransactionsHandler creates a handler for listing transactions
func NewListTransactionsHandler(usecase adminUsecase.ListTransactionsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "100")
		offsetStr := c.DefaultQuery("offset", "0")

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 100
		}
		if limit > 1000 {
			limit = 1000
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			offset = 0
		}

		output, appErr := usecase.Execute(c, limit, offset)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
