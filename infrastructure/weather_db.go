package infrastructure

import (

	//"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/go-sql-driver/mysql"
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* var db sql.DB
 */
type WeatherDB struct {
	//*sql.DB
	*gorm.DB
}

func NewWeatherDb(config DBConfig) (*WeatherDB, error) {
	var err error

	dsn := config.FormatDSN()
	sqlDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	db, _ := sqlDb.DB()
	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.MaxConns)
	db.SetConnMaxLifetime(config.MaxLifeTIme)
	sqlDb.WithContext(context.Background())
	/* 	if db, err := sql.Open("mysql", dbConfig.FormatDSN()); err == nil {
		db.SetMaxOpenConns(config.MaxConns)
		db.SetMaxIdleConns(config.MaxConns)
		db.SetConnMaxLifetime(config.MaxLifeTIme)
		return &WeatherDB{
			db,
		}, nil
	} */

	return &WeatherDB{
		sqlDb,
	}, nil
}

func (db *WeatherDB) Close() {
	db.Close()
}

/* func (db *DB) Open(config DBConfig) {

} */
