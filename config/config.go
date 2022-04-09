package config

import (
	"flag"
	"fmt"
)

type Config struct {
	HttpAddr string
	ApiKey   string

	DB *DatabaseConfig

	ApiVK *VKConfig
}

type ServiceConfig struct {
	Host string
}

type DatabaseConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func (d DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%s:%s)/%v?charset=utf8mb4&parseTime=true&loc=Local", d.User, d.Pass, d.Host, d.Port, d.Name)
}

type VKConfig struct {
	ApiKey string
}

// GetConfig - функция получения конфигурационных настроек
func GetConfig() Config {
	cfg := Config{}
	cfg.DB = &DatabaseConfig{}
	cfg.ApiVK = &VKConfig{}

	flag.StringVar(&cfg.HttpAddr, "HTTP_ADDR", ":8080", "")
	flag.StringVar(&cfg.ApiKey, "API_KEY", "", "")

	flag.StringVar(&cfg.ApiVK.ApiKey, "VK_API_KEY", "", "")

	flag.StringVar(&cfg.DB.Host, "DB_HOST", "", "")
	flag.StringVar(&cfg.DB.Port, "DB_PORT", "", "")
	flag.StringVar(&cfg.DB.Name, "DB_NAME", "", "")
	flag.StringVar(&cfg.DB.User, "DB_USER", "", "")
	flag.StringVar(&cfg.DB.Pass, "DB_PASS", "", "")

	flag.Parse()

	return cfg
}
