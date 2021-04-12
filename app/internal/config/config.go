package config

type AppConfig struct {
	Server `json:"server"`
}

type Server struct {
	Addr string `json:"addr"`
}
