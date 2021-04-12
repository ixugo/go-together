package global

import (
	"github.com/spf13/viper"
	"together/app/internal/config"
)

var (
	AppConfig *config.AppConfig
	Viper     *viper.Viper
)
