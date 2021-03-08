package infrastructure

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var configObject *Config = &Config{}

func init() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs")

	// read from env variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	configObject = &Config{}
	if err := viper.Unmarshal(configObject); err != nil {
		panic(err)
	}

	fmt.Println(configObject)
}

type Config struct {
	Crawler CrawlerConfig `mapstructure:"crawler"`
	DB      DBConfig      `mapstructure:"db"`
	Secret  string        `mapstructure:"secret"`
}

type CrawlerConfig struct {
	Interval string `mapstructure:"interval"`
	CwbUrl   string `mapstructure:"cwb_url"`
	Token    string `mapstructure:"token"`
}

type DBConfig struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	User        string        `mapstructure:"user"`
	Password    string        `mapstructure:"password"`
	MaxConns    int           `mapstructure:"max_conns"`
	MaxLifeTIme time.Duration `mapstructure:"max_lifetime"`
}

func (dbconfig DBConfig) FormatDSN() string {
	dbConfigObject := &mysql.Config{
		User:                 dbconfig.User,
		Passwd:               dbconfig.Password,
		Addr:                 dbconfig.Host,
		Net:                  "tcp",
		DBName:               "weather_forecast",
		AllowNativePasswords: true,
	}
	return dbConfigObject.FormatDSN()
}

func GetConfig() *Config {
	return configObject
}

func GetSecret() string {
	return configObject.Secret
}
