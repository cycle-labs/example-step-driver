package driver

import (
	"github.com/cycle-labs/example-step-driver/starter/api"
	"github.com/gin-gonic/gin"
)

type GinHandlers struct {
	service *DriverService
}

func NewGinHandlers(service *DriverService) *GinHandlers {
	return &GinHandlers{service: service}
}

// Retrieve account balance
// (GET /get_balance)
func (h *GinHandlers) GetGetBalance(c *gin.Context, params api.GetGetBalanceParams) {
	// TODO: Provide Implementation
	panic("not implemented")
}

// Transfer funds between accounts
// (POST /transfer)
func (h *GinHandlers) PostTransfer(c *gin.Context) {
	// TODO: Provide Implementation
	panic("not implemented")
}
