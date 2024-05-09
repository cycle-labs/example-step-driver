package driver

import (
	"net/http"

	"github.com/cycle-labs/example-step-driver/sp-driver-stateless/api"
	"github.com/gin-gonic/gin"
)

type GinHandlers struct {
	service *DriverService
}

func NewGinHandlers(service *DriverService) *GinHandlers {
	return &GinHandlers{service: service}
}

// Return the account balance
// (GET /balance/{accountID})
func (h *GinHandlers) GetBalance(c *gin.Context, accountID int) {
	b, err := h.service.GetBalance(c, accountID)
	if err != nil {
		errMessage := err.Error()
		c.JSON(http.StatusOK, api.StepResponse{
			Status:       api.Fail,
			ErrorMessage: &errMessage,
		})
		return
	}
	balance := &api.StepResponse_Variables_AdditionalProperties{}
	balance.FromStepResponseVariables0(b)
	variables := make(map[string]api.StepResponse_Variables_AdditionalProperties)
	variables["balance"] = *balance
	c.JSON(http.StatusOK, api.StepResponse{
		Status:    api.Pass,
		Variables: &variables,
	})
}

// Execute a transfer
// (POST /transfer)
func (h *GinHandlers) Transfer(c *gin.Context) {
	var t api.Transfer
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	amount, err := FromMonetaryValue(t.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorMessage{Message: err.Error()})
		return
	}
	err = h.service.Transfer(c, t.Sender, t.Receiver, *amount)
	if err != nil {
		errMessage := err.Error()
		c.JSON(http.StatusOK, api.StepResponse{
			Status:       api.Fail,
			ErrorMessage: &errMessage,
		})
		return
	}
	c.JSON(http.StatusOK, api.StepResponse{Status: api.Pass})
}
