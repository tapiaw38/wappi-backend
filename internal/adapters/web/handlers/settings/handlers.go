package settings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
	settingsUsecase "yego/internal/usecases/settings"
)

// NewGetHandler creates a handler for getting settings
func NewGetHandler(usecase settingsUsecase.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, appErr := usecase.Execute(c)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output.Settings)
	}
}

type UpdateInput struct {
	BusinessName        *string  `json:"business_name,omitempty"`
	BusinessLatitude    *float64 `json:"business_latitude,omitempty"`
	BusinessLongitude   *float64 `json:"business_longitude,omitempty"`
	DefaultMapLatitude  *float64 `json:"default_map_latitude,omitempty"`
	DefaultMapLongitude *float64 `json:"default_map_longitude,omitempty"`
	DefaultMapZoom      *int     `json:"default_map_zoom,omitempty"`
	DefaultItemWeight   *int     `json:"default_item_weight,omitempty"`
	DeliveryBasePrice   *float64 `json:"delivery_base_price,omitempty"`
	DeliveryPricePerKm  *float64 `json:"delivery_price_per_km,omitempty"`
	DeliveryPricePerKg  *float64 `json:"delivery_price_per_kg,omitempty"`
	ManagerCollectorID  *string  `json:"manager_collector_id,omitempty"`
}

// NewUpdateHandler creates a handler for updating settings
func NewUpdateHandler(usecase settingsUsecase.UpdateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		output, appErr := usecase.Execute(c, settingsUsecase.UpdateInput{
			BusinessName:       input.BusinessName,
			BusinessLatitude:   input.BusinessLatitude,
			BusinessLongitude:  input.BusinessLongitude,
			DefaultMapLatitude: input.DefaultMapLatitude,
			DefaultMapLongitude: input.DefaultMapLongitude,
			DefaultMapZoom:     input.DefaultMapZoom,
			DefaultItemWeight:  input.DefaultItemWeight,
			DeliveryBasePrice:  input.DeliveryBasePrice,
			DeliveryPricePerKm: input.DeliveryPricePerKm,
			DeliveryPricePerKg: input.DeliveryPricePerKg,
			ManagerCollectorID: input.ManagerCollectorID,
		})
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output.Settings)
	}
}

type CalculateDeliveryFeeItemInput struct {
	Quantity int  `json:"quantity"`
	Weight   *int `json:"weight,omitempty"`
}

type CalculateDeliveryFeeInput struct {
	UserLatitude  float64                         `json:"user_latitude" binding:"required"`
	UserLongitude float64                         `json:"user_longitude" binding:"required"`
	Items         []CalculateDeliveryFeeItemInput `json:"items"`
}

// NewCalculateDeliveryFeeHandler creates a handler for calculating delivery fee
func NewCalculateDeliveryFeeHandler(usecase settingsUsecase.CalculateDeliveryFeeUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CalculateDeliveryFeeInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewApplicationError(mappings.RequestBodyParsingError, err)
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		// Map handler input to usecase input
		usecaseInput := settingsUsecase.CalculateDeliveryFeeInput{
			UserLatitude:  input.UserLatitude,
			UserLongitude: input.UserLongitude,
		}

		// Convert items
		for _, item := range input.Items {
			usecaseInput.Items = append(usecaseInput.Items, struct {
				Quantity int  `json:"quantity"`
				Weight   *int `json:"weight,omitempty"`
			}{
				Quantity: item.Quantity,
				Weight:   item.Weight,
			})
		}

		output, appErr := usecase.Execute(c, usecaseInput)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
