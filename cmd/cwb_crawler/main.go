package main

import (
	"os"
	"os/signal"
	"syscall"
	"weather_forecast/application/cwb_crawler"
	"weather_forecast/infrastructure"
	"weather_forecast/repository"
)

func main() {
	db, err := infrastructure.NewWeatherDb(infrastructure.GetConfig().DB)
	if err != nil {
		panic(err)
	}
	repo := repository.NewForecastRepository(db)
	app := cwb_crawler.NewCrawlerApplication(repo, infrastructure.GetConfig().Crawler)
	app.PrepareJobContext()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	defer db.Close()
}
