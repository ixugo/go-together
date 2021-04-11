package global

import (
	"github.com/spf13/viper"
	"together/app/internal/config"
)

var (
	APP_CONFIG *config.AppConfig
	VIPER      *viper.Viper
)
