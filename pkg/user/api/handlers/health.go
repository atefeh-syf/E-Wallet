package handlers

import (
	"net/http"

	"github.com/atefeh-syf/yumigo/pkg/user/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))
}