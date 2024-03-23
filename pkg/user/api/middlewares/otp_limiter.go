package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/atefeh-syf/yumigo/pkg/user/api/helper"
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/atefeh-syf/yumigo/pkg/user/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("Not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}