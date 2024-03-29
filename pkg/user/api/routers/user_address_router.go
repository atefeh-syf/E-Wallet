package routers

import (
	"github.com/atefeh-syf/yumigo/pkg/user/api/handlers"
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/gin-gonic/gin"
)

func UserAddress(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUserAddressHandler(cfg)

	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
	router.POST("/get-by-filter", h.GetByFilter)
}
