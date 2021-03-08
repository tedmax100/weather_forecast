package cwb_crawler

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"weather_forecast/domain/cwb_crawl/model"
	"weather_forecast/infrastructure"
	"weather_forecast/repository"

	"github.com/robfig/cron/v3"
)

type CrawlerApplication struct {
	url        string
	interval   string
	repository *repository.ForecastRepository
	cron       *cron.Cron
	jobs       map[cron.EntryID]Job
}

func NewCrawlerApplication(repo *repository.ForecastRepository, config infrastructure.CrawlerConfig) *CrawlerApplication {
	return &CrawlerApplication{
		url:        fmt.Sprintf(config.CwbUrl, config.Token),
		interval:   config.Interval,
		repository: repo,
		cron:       cron.New(),
		jobs:       make(map[cron.EntryID]Job),
	}
}

func (c *CrawlerApplication) PrepareJobContext() error {
	// 取得location list
	// 設定排程任務
	//var locations []model.Location
	var err error
	var elements []model.Elemant
	var locations []model.Location
	if locations, err = c.repository.GetLocations(context.Background()); err != nil {
		return err
	}
	c.appendLocationToUrl(locations)

	if elements, err = c.repository.GetElementss(context.Background()); err != nil {
		return err
	}
	// prepare meta data finish
	// prepare job
	job := NewCwbCrawlerJob(c.url, locations, elements, c.repository)
	if jobId, err := c.cron.AddJob(c.interval, job); err == nil {
		c.jobs[jobId] = job
	} else {
		log.Fatal(err)
	}

	c.cron.Start()

	// initial round to run the job
	job.Run()
	return nil

}

func (c *CrawlerApplication) Start() {

}

func (c *CrawlerApplication) Stop() {

}

func (c *CrawlerApplication) locationNames(locations []model.Location) []string {
	var locationNames []string = make([]string, len(locations))
	for idx, location := range locations {
		locationNames[idx] = location.Name
	}
	return locationNames
}

func (c *CrawlerApplication) appendLocationToUrl(locations []model.Location) {
	v := url.Values{}
	v.Add("locationName", strings.Join(c.locationNames(locations), ","))
	c.url += "&" + v.Encode()
}
