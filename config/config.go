package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

const (
	// Dev develop environment
	Dev string = "dev"
	// Test test environment
	Test string = "test"
	// Pre preview environment
	Pre string = "pre"
	// Prod production environment
	Prod string = "prod"
)

// Items Items
var Items = func() configuration {
	return config
}
var config configuration

type configuration struct {
	AppName    string `env:"APP_NAME" default:"mweather-go"`
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	APIVersion string `env:"API_VERSION" default:"Commit ID"`
	Mysql      struct {
		Host    string `default:"mysql-sz.makeblock.com"`
		Port    string `default:"3306"`
		DBName  string `default:"mweather-go"`
		User    string `default:"root"`
		Pwd     string `default:"tEKIZtC0CXjg"`
		Charset string `default:"utf8mb4"`
	}
	Redis struct {
		Host   string `default:"127.0.0.1"`
		Port   string `default:"6379"`
		Pwd    string `default:""`
		Prefix string `default:"'mweather-go:'"`
	}
	RPC struct {
		Account string `default:"account-service-dev.makeblock.com:8000"` //RPC_ACCOUNT
	}
}

// Load load configurations
func Load() {
	conf := &configor.Config{ENVPrefix: "-"}
	if err := configor.New(conf).Load(&config, "/mweather-go/etc/config.yaml", "config.yaml", "../config.yaml"); err != nil {
		panic(err)
	}
	fmt.Printf("config: %+v\n", config)
}
