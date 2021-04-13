package configs

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Engine struct {
	v *viper.Viper
}

func LoadConfig(configs ...string) *Engine {
	v := viper.New()
	v.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			v.AddConfigPath(config)
		}
	}
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("请检查配置文件 err:%w", err))
	}

	s := Engine{v}
	s.WatchConfig()
	return &s
}

func (s *Engine) WatchConfig() {
	s.v.WatchConfig()
	s.v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("onChange the config")
		_ = s.ReloadConfig()
	})
}

var engine = make(map[string]interface{})

func (s *Engine) Read(k string, v interface{}) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := engine[k]; !ok {
		engine[k] = v
	}
	return nil
}

func (s *Engine) ReloadConfig() error {
	for k, v := range engine {
		err := s.Read(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
