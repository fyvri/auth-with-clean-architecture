package auth

import (
	"auth-with-clean-architecture/modules/auth/entity"
	user_entity "auth-with-clean-architecture/modules/user/entity"
	"auth-with-clean-architecture/pkg/password"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UseCase struct {
	r RepositoryInterface
}
type UseCaseInterface interface {
	Login(payload *entity.Payload) (*user_entity.User, string, error)
	ShowProfile(tokenSigned string) (*user_entity.User, error)
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		r: r,
	}
}

func (uc *UseCase) Login(payload *entity.Payload) (*user_entity.User, string, error) {
	user, _ := uc.r.FindByUsername(payload.Username)
	if user.Username == "" {
		return nil, "", errors.New("user not found")
	}

	match := password.CheckPasswordHash(payload.Password, user.Password)
	if !match {
		return nil, "", errors.New("password is incorrect")
	}

	expTime := time.Now().Add(time.Minute * 3600)
	claims := &entity.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-token",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(entity.JWT_KEY)
	if err != nil {
		return nil, "", err
	}

	return user, tokenSigned, err
}

func (uc *UseCase) ShowProfile(tokenSigned string) (*user_entity.User, error) {
	token, err := jwt.Parse(tokenSigned, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(entity.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	res, err := uc.r.FindByUsername(claims["Username"].(string))
	if res == nil {
		return nil, fmt.Errorf("user not found")
	}

	return res, err
}
