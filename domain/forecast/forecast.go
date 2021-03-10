package forecast

import (
	"context"
	"log"
	"weather_forecast/domain/forecast/model"
	"weather_forecast/repository"
)

type ForecastDomain struct {
	repo *repository.ForecastRepository
}

func (d *ForecastDomain) Forecast(contextData model.ForecastRequest) (model.Records, error) {
	var result model.Records = model.Records{}
	nearestLastStartTime, nearestLastEndTime := contextData.GetNearestLastPeriod()
	records, err := d.repo.GetRecords(context.Background(), model.ForecastQueryDao{StartTime: nearestLastStartTime, EndTime: nearestLastEndTime})
	if err != nil {
		return result, err
	}

	var preLocationName string
	var locations []model.Location = make([]model.Location, 0)
	var preLocationIdx = 0
	for _, data := range records {
		if preLocationName != data.Location {
			location := model.Location{LocationName: data.Location}
			location.WeatherElement = make([]model.WeatherElement, 0)
			locations = append(locations, location)
			preLocationName = data.Location
			preLocationIdx++
		}

		element := model.WeatherElement{
			ElementName: data.Element,
			Time: []model.TimeRecord{
				model.TimeRecord{
					StartTime: data.StartTime,
					EndTime:   data.EndTime,
					Parameter: model.Parameter{
						ParameterName:  data.ParameterName,
						ParameterUnit:  data.ParameterUnit,
						ParameterValue: data.ParameterValue,
					},
				},
			},
		}
		locations[preLocationIdx-1].WeatherElement = append(locations[preLocationIdx-1].WeatherElement, element)
	}
	log.Println(contextData.GetNearestLastPeriod())
	result.Location = locations
	return result, err
}

func NewForecastDomain(repo *repository.ForecastRepository) *ForecastDomain {
	return &ForecastDomain{
		repo: repo,
	}
}
