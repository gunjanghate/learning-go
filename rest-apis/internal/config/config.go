package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// what is struct tags
// Struct tags are metadata added to struct fields in Go. They provide additional information about the field and can be used by various libraries and frameworks for tasks such as serialization, validation, and configuration.
// For example, the `yaml:"env"` tag indicates that the Env field corresponds to the "env" key in a YAML file, while the `env:"ENV"` tag indicates that it can also be set using the "ENV" environment variable. The `env-required:"true"` tag specifies that this field is mandatory and must be provided either through the YAML file or an environment variable.

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" ` // struct tags for yaml and env
	StoragePath string `yaml:"storage_path" env-required:"true" `
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config","", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config file does not exist at path: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	return &cfg
}