package repository

import (
	"context"
	"log"
	"weather_forecast/domain/cwb_crawl/model"
	forecastModel "weather_forecast/domain/forecast/model"
	userModel "weather_forecast/domain/user/model"
	"weather_forecast/infrastructure"

	"gorm.io/gorm/clause"
)

type ForecastRepository struct {
	forecastDb *infrastructure.WeatherDB
}

func NewForecastRepository(weatherDb *infrastructure.WeatherDB) *ForecastRepository {
	return &ForecastRepository{
		forecastDb: weatherDb,
	}
}

func (l *ForecastRepository) GetLocations(ctx context.Context) ([]model.Location, error) {
	var locations []model.Location = make([]model.Location, 0)
	result := l.forecastDb.Table("location").Find(&locations)
	if result.Error != nil {
		return locations, result.Error
	}
	return locations, nil
}

func (l *ForecastRepository) GetElementss(ctx context.Context) ([]model.Elemant, error) {
	var elements []model.Elemant = make([]model.Elemant, 0)
	result := l.forecastDb.Table("element").Find(&elements)
	if result.Error != nil {
		return elements, result.Error
	}
	return elements, nil
}

func (l *ForecastRepository) CreateRecords(ctx context.Context, records []model.ForcasstDao) error {
	result := l.forecastDb.Table("record").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "location_id"}, {Name: "element_id"}, {Name: "start_time"}},
		DoUpdates: clause.AssignmentColumns([]string{"parameter_name", "parameter_value"}),
	}).Create(records)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r *ForecastRepository) GetUser(ctx context.Context, user userModel.User) (userModel.User, error) {
	var data userModel.User = userModel.User{}
	result := r.forecastDb.Table("user").Where(&userModel.User{Name: user.Name}).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (r *ForecastRepository) GetRecords(ctx context.Context, queryContext forecastModel.ForecastQueryDao) ([]*forecastModel.RecordData, error) {
	var data []*forecastModel.RecordData = make([]*forecastModel.RecordData, 0)
	log.Println(queryContext.StartTime.String())
	result := r.forecastDb.Debug().Raw(`SELECT l.name AS location, e.name AS element, r.parameter_name, r.parameter_unit, r.parameter_value, r.start_time, r.end_time 
	FROM record AS r
	INNER join element AS e ON r.element_id = e.id
	INNER join location AS l ON r.location_id = l.id
	WHERE r.start_time = ? ORDER BY l.id;`, queryContext.StartTime.Format("2006-01-02 15:04:05")).Scan(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil

}
