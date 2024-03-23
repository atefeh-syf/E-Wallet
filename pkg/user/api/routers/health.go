package routers

import (
	"github.com/atefeh-syf/yumigo/pkg/user/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}