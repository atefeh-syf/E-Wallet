package middlewares

import (
	"context"
	"strings"

	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/atefeh-syf/yumigo/pkg/user/pkg/service_errors"
	"github.com/atefeh-syf/yumigo/pkg/user/services"
	"github.com/golang-jwt/jwt"
)

func Authentication(authToken string, c context.Context) (error , map[string]interface{}) {
	cfg := config.GetConfig()
	tokenService := services.NewTokenService(cfg)
	

	claimMap := map[string]interface{}{}
	token := strings.Split(authToken, " ")
	var err error
	if authToken == "" {
		err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
	} else {
		claimMap, err = tokenService.GetClaims(token[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
			default:
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
			}
		}
	}
	if err != nil {
		return err, map[string]interface{}{}
	}
	return nil, claimMap
}
