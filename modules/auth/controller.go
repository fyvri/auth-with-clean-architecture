package auth

import (
	"auth-with-clean-architecture/modules/auth/entity"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	Login(body *AuthRequest) (*UserItemAndToken, error)
	ShowProfile(tokenSigned string) (*UserItem, error)
	VerifyToken(tokenSigned string) (*entity.JWTClaim, error)
}

func NewController(uc UseCaseInterface) ControllerInterface {
	return &Controller{
		uc: uc,
	}
}

type UserItem struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
}

type UserItemAndToken struct {
	User  *UserItem `json:"user"`
	Token string    `json:"token"`
}

func (c *Controller) Login(body *AuthRequest) (*UserItemAndToken, error) {
	payload := entity.Payload{
		Username: body.Username,
		Password: body.Password,
	}
	user, token, err := c.uc.Login(&payload)
	if err != nil {
		return nil, err
	}

	res := &UserItemAndToken{
		User: &UserItem{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			RoleID:   user.RoleID,
		},
		Token: token,
	}

	return res, nil
}

func (c *Controller) ShowProfile(tokenSigned string) (*UserItem, error) {
	user, err := c.uc.ShowProfile(tokenSigned)
	if err != nil {
		return nil, err
	}

	return &UserItem{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		RoleID:   user.RoleID,
	}, nil
}

func (c *Controller) VerifyToken(tokenSigned string) (*entity.JWTClaim, error) {
	token, err := jwt.Parse(tokenSigned, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(entity.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)

		return &entity.JWTClaim{
			Username: claims["Username"].(string),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "jwt-token",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3600)),
			},
		}, nil
	} else {
		return nil, fmt.Errorf("invalid authorization token")
	}
}
