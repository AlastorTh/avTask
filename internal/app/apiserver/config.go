package apiserver

import "github.com/AlastorTh/avTask/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `yaml:"bind_addr"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Store:    store.NewConfig(),
	}
}
