package order

import (
	"net/http"

	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	orderUsecase "yego/internal/usecases/order"

	"github.com/gin-gonic/gin"
)

type CreateWithLinkItemInput struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Weight   *int    `json:"weight,omitempty"`
}

type CreateWithLinkDataInput struct {
	Items []CreateWithLinkItemInput `json:"items"`
}

type CreateWithLinkInput struct {
	PhoneNumber string                   `json:"phone_number"`
	ETA         string                   `json:"eta"`
	Data        *CreateWithLinkDataInput `json:"data,omitempty"`
}

// NewCreateWithLinkHandler creates a handler for creating orders with claim links
// This endpoint is intended to be called by the WhatsApp IA
func NewCreateWithLinkHandler(usecase orderUsecase.CreateWithLinkUsecase, frontendURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateWithLinkInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		// Map handler input to usecase input
		usecaseInput := orderUsecase.CreateWithLinkInput{
			PhoneNumber: input.PhoneNumber,
			ETA:         input.ETA,
		}

		if input.Data != nil && len(input.Data.Items) > 0 {
			items := make([]orderUsecase.CreateWithLinkItemInput, len(input.Data.Items))
			for i, item := range input.Data.Items {
				items[i] = orderUsecase.CreateWithLinkItemInput{
					Name:     item.Name,
					Price:    item.Price,
					Quantity: item.Quantity,
					Weight:   item.Weight,
				}
			}
			usecaseInput.Data = &orderUsecase.CreateWithLinkDataInput{Items: items}
		}

		output, appErr := usecase.Execute(c, usecaseInput, frontendURL)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
