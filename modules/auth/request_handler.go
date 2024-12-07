package auth

import (
	"auth-with-clean-architecture/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	c ControllerInterface
}

type RequestHandlerInterface interface {
	Login(c *gin.Context)
	ShowProfile(c *gin.Context)
}

func NewRequestHandler(c ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		c: c,
	}
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (rh *RequestHandler) Login(c *gin.Context) {
	var req AuthRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Code:    500,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	data, err := rh.c.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Code:    500,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Code:    200,
			Message: "login successfully",
		},
		Data: data,
	})
}

func (rh *RequestHandler) ShowProfile(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if authorization == nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Code:    401,
				Message: "Unauthorized",
			},
			Data: nil,
		})
		return
	}

	tokenSigned := strings.Split(authorization[0], " ")[1]
	res, err := rh.c.ShowProfile(tokenSigned)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Code:    500,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Code:    200,
			Message: "",
		},
		Data: res,
	})
}
