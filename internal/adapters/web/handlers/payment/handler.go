package payment

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"wappi/internal/platform/appcontext"
)

type Handler struct {
	contextFactory appcontext.Factory
}

func NewHandler(contextFactory appcontext.Factory) *Handler {
	return &Handler{
		contextFactory: contextFactory,
	}
}

func (h *Handler) CheckPaymentMethod(c *gin.Context) {
	username := c.Param("user_id")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	app := h.contextFactory()

	// Resolve the internal UUID from the username via auth-api
	authHeader := c.GetHeader("Authorization")
	var token string
	if authHeader != "" && len(authHeader) > 7 {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	internalUserID, err := app.Integrations.Auth.GetUserIDByUsername(username, token)
	if err != nil {
		log.Printf("Warning: could not resolve user ID for username %s: %v", username, err)
		internalUserID = username
	}

	hasPaymentMethod, err := app.Integrations.Payments.HasPaymentMethod(internalUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check payment method"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_payment_method": hasPaymentMethod})
}
