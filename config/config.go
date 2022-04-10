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
	ConfirmationKey string
	CallbackUrl     string
	ApiKey          string
}

// GetConfig - функция получения конфигурационных настроек
func GetConfig() Config {
	cfg := Config{}
	cfg.DB = &DatabaseConfig{}
	cfg.ApiVK = &VKConfig{}

	flag.StringVar(&cfg.HttpAddr, "HTTP_ADDR", ":8080", "")
	flag.StringVar(&cfg.ApiKey, "API_KEY", "0abdbfa7-b7f0-4673-b9bb-ab826fdd4d1f", "")

	flag.StringVar(&cfg.ApiVK.CallbackUrl, "CALLBACK_URL", "http://194.58.120.99/callback", "")
	flag.StringVar(&cfg.ApiVK.ConfirmationKey, "CONFIRMATION_KEY", "bc90fc38", "")
	flag.StringVar(&cfg.ApiVK.ApiKey, "VK_API_KEY", "445c6e92d5370fc4c486f11f8d3b088c96fa795314f1b5432dcce5c63da1cf42afb4ad100a1b2b7047578", "")

	flag.StringVar(&cfg.DB.Host, "DB_HOST", "localhost", "")
	flag.StringVar(&cfg.DB.Port, "DB_PORT", "10000", "")
	flag.StringVar(&cfg.DB.Name, "DB_NAME", "chat_bot", "")
	flag.StringVar(&cfg.DB.User, "DB_USER", "user", "")
	flag.StringVar(&cfg.DB.Pass, "DB_PASS", "user", "")

	flag.Parse()

	return cfg
}
