package routers

import (
	"github.com/atefeh-syf/yumigo/pkg/user/api/handlers"
	"github.com/atefeh-syf/yumigo/pkg/user/api/middlewares"
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}