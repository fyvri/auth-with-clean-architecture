package middleware

import (
	"auth-with-clean-architecture/dto"
	"auth-with-clean-architecture/modules/auth"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader(authorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Code:    401,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Code:    401,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Code:    401,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	accessToken := fields[1]
	payload, err := auth.ControllerInterface.VerifyToken(&auth.Controller{}, accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Code:    401,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.Set(authorizationPayloadKey, payload)
	c.Next()
}
