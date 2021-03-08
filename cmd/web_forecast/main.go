package main

import (
	"weather_forecast/application/web_forecast"
	"weather_forecast/infrastructure"
	"weather_forecast/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := infrastructure.NewWeatherDb(infrastructure.GetConfig().DB)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	repo := repository.NewForecastRepository(db)

	engine := gin.Default()
	web_forecast.InitRouter(engine, repo)
	engine.Run()
}
