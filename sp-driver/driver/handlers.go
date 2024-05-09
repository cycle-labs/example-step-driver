package driver

import (
	"net/http"

	"github.com/cycle-labs/example-step-driver/sp-driver/api"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type GinHandlers struct {
	service *DriverService
}

func NewGinHandlers(service *DriverService) *GinHandlers {
	return &GinHandlers{service: service}
}

// Create a new session
// (POST /sessions)
func (h *GinHandlers) Login(c *gin.Context) {
	var login api.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateSession(c, login.Username, login.Password, login.Url)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.ErrorMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, api.LoginResponse{SessionID: id})
}

// End the session
// (DELETE /sessions/{sessionID})
func (h *GinHandlers) Logout(c *gin.Context, sessionID openapi_types.UUID) {
	err := h.service.CloseSession(c, sessionID)
	if err != nil {
		if err == ErrSessionNotFound {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, api.ErrorMessage{Message: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Return the account balance
// (GET /sessions/{sessionID}/balance/{accountID})
func (h *GinHandlers) GetBalance(c *gin.Context, sessionID openapi_types.UUID, accountID int) {
	b, err := h.service.GetBalance(c, sessionID, accountID)
	if err != nil {
		if err == ErrSessionNotFound {
			c.Status(http.StatusNotFound)
			return
		}
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
// (POST /sessions/{sessionID}/transfer)
func (h *GinHandlers) Transfer(c *gin.Context, sessionID openapi_types.UUID) {
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
	transfer := Transfer{
		FromAccountID: t.Sender,
		ToAccountID:   t.Receiver,
		Amount:        *amount,
	}
	err = h.service.Transfer(c, sessionID, transfer)
	if err != nil {
		if err == ErrSessionNotFound {
			c.Status(http.StatusNotFound)
			return
		}
		errMessage := err.Error()
		c.JSON(http.StatusOK, api.StepResponse{
			Status:       api.Fail,
			ErrorMessage: &errMessage,
		})
		return
	}
	c.JSON(http.StatusOK, api.StepResponse{Status: api.Pass})
}
