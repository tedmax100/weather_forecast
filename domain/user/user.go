package user

import (
	"context"
	"strconv"
	"time"
	"weather_forecast/domain/user/model"
	"weather_forecast/infrastructure"
	"weather_forecast/repository"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(infrastructure.GetSecret())

type UserDomain struct {
	repo *repository.ForecastRepository
}

func NewUserDomain(repo *repository.ForecastRepository) *UserDomain {
	return &UserDomain{
		repo: repo,
	}
}

func (u *UserDomain) Login(contextData model.UserLoginRequest) (model.UserLoginResponse, error) {
	userData, _ := u.repo.GetUser(context.Background(), model.User{Name: contextData.Name})

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(contextData.Password)); err != nil {
		return model.UserLoginResponse{}, err
	}

	now := time.Now()
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(1 * time.Hour).Unix(),
		Issuer:    "forecast",
		IssuedAt:  now.Unix(),
		Audience:  userData.Name,
		Id:        userData.Name + strconv.FormatInt(now.Unix(), 10),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return model.UserLoginResponse{Token: token}, err
}
