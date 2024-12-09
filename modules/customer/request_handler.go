package customer

import (
	"auth-with-clean-architecture/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	c ControllerInterface
}

type RequestHandlerInterface interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewRequestHandler(c ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		c: c,
	}
}

type CreateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

func (rh *RequestHandler) List(c *gin.Context) {
	res, err := rh.c.List()
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
			Code: 500,
		},
		Data: res,
	})
}

func (rh *RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

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

	res, err := rh.c.Create(&req)
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
			Message: "customer successfully created",
		},
		Data: res,
	})
}

func (rh *RequestHandler) Read(c *gin.Context) {
	ID := c.Param("ID")
	res, err := rh.c.Read(ID)
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
			Code: 200,
		},
		Data: res,
	})
}

func (rh *RequestHandler) Update(c *gin.Context) {
	req := CreateRequest{}
	ID := c.Param("ID")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Code:    500,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := rh.c.Update(ID, &req)
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
			Message: "customer successfully updated",
		},
		Data: res,
	})
}

func (rh *RequestHandler) Delete(c *gin.Context) {
	ID := c.Param("ID")
	err := rh.c.Delete(ID)
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
			Message: "customer successfully deleted",
		},
	})
}
