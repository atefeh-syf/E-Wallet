package handlers

import (
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/atefeh-syf/yumigo/pkg/user/services"
	"github.com/gin-gonic/gin"
)

type UserAddressHandler struct {
	service *services.UserAddressService
}

func NewUserAddressHandler(cfg *config.Config) *UserAddressHandler {
	service := services.NewUserAddressService(cfg)
	return &UserAddressHandler{service: service}
}

func (h *UserAddressHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *UserAddressHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *UserAddressHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *UserAddressHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *UserAddressHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}