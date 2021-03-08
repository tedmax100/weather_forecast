package web_forecast

import (
	"log"
	"net/http"
	"time"
	"weather_forecast/domain/forecast"
	forecastModel "weather_forecast/domain/forecast/model"
	"weather_forecast/domain/user"
	"weather_forecast/domain/user/model"
	"weather_forecast/middleware"
	"weather_forecast/repository"

	"github.com/gin-gonic/gin"
)

var userDomain *user.UserDomain
var forecastDomain *forecast.ForecastDomain

func InitRouter(engine *gin.Engine, repo *repository.ForecastRepository) {
	userDomain = user.NewUserDomain(repo)
	forecastDomain = forecast.NewForecastDomain(repo)

	engine.POST("/login", Login)

	forecastGrp := engine.Group("/forecast")
	forecastGrp.Use(middleware.AuthJwt())
	forecastGrp.GET("/now", ForecastNow)
}

func Login(c *gin.Context) {
	var userRequest model.UserLoginRequest
	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if userRequest.IsValid() == false {
		c.Status(http.StatusBadRequest)
		return
	}

	res, err := userDomain.Login(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, res)
	return
}

func ForecastNow(c *gin.Context) {
	// user :=	c.GetString("user")
	var forecastRequest forecastModel.ForecastRequest = forecastModel.ForecastRequest{
		CurrentTime: time.Now(),
	}
	locationName := c.Query("locationName")
	if locationName != "" {
		forecastRequest.Location = locationName
	}

	log.Println(forecastRequest)
	records, err := forecastDomain.Forecast(forecastRequest)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		return
	}
	c.JSON(http.StatusOK, records)
	return
}
