package cwb_crawler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"weather_forecast/domain/cwb_crawl/model"
	"weather_forecast/repository"
)

type Job interface {
	Run()
}

type CwbCrawlerJob struct {
	//locations  []model.Location
	url                string
	repository         *repository.ForecastRepository
	locationDictionary map[string]uint //[]model.Location
	elementDictionary  map[string]uint //[]model.Elemant
}

func NewCwbCrawlerJob(url string, locations []model.Location, elements []model.Elemant, repository *repository.ForecastRepository) *CwbCrawlerJob {
	var locationDictionary map[string]uint = make(map[string]uint)
	var elementDictionary map[string]uint = make(map[string]uint)
	for _, location := range locations {
		locationDictionary[location.Name] = location.Id
	}

	for _, element := range elements {
		elementDictionary[element.Name] = element.Id
	}

	return &CwbCrawlerJob{
		url:                url,
		repository:         repository,
		locationDictionary: locationDictionary,
		elementDictionary:  elementDictionary,
	}
}

func (ccj *CwbCrawlerJob) Run() {
	res, err := http.Get(ccj.url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	weatherData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var cwbData model.CwbResponseData
	if err := json.Unmarshal(weatherData, &cwbData); err != nil {
		log.Fatal(err)
	}

	var forecastDaos []model.ForcasstDao = make([]model.ForcasstDao, 0)

	for _, location := range cwbData.Records.Location {
		for _, elementData := range location.WeatherElement {
			for _, forecastData := range elementData.Time {
				log.Println(forecastData)
				var forecastDao model.ForcasstDao = model.ForcasstDao{}
				forecastDao.LocationId = ccj.locationDictionary[location.LocationName]
				forecastDao.ElementId = ccj.elementDictionary[elementData.ElementName]
				forecastDao.StartTime, _ = time.Parse("2006-01-02 15:04:05", forecastData.StartTime)
				forecastDao.EndTime, _ = time.Parse("2006-01-02 15:04:05", forecastData.EndTime)
				forecastDao.ParameterUnit = forecastData.Parameter.ParameterUnit
				forecastDao.ParameterName = forecastData.Parameter.ParameterName
				if forecastData.Parameter.ParameterValue != nil {
					parameterValue, _ := strconv.ParseInt(*forecastData.Parameter.ParameterValue, 10, 64)
					forecastDao.ParameterValue = new(int)
					*forecastDao.ParameterValue = int(parameterValue)
				}

				forecastDaos = append(forecastDaos, forecastDao)
			}
		}
	}
	ccj.repository.CreateRecords(context.TODO(), forecastDaos)

}
