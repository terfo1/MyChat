package configs

import (
	"github.com/terfo1/news/internal/templates"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName string `yaml:"app_name"`
	Prefork bool   `yaml:"prefork"`
	Port    string `yaml:"port"`
}

var AppConfig *Config

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Couldn't unmarshal config: %v", err)
	}

	AppConfig = &config
	return AppConfig
}

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName: AppConfig.AppName,
		Prefork: AppConfig.Prefork,
		Views:   templates.TemplateEngine,
	}
}
