package configs

import (
	"fmt"
	//"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerPort string `yaml:"server_port"`
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

func LoadConfig() (*Config, error) {
	// Cargar las configuraciones del archivo YAML
	configFile, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo de configuraciones: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("error al decodificar el archivo de configuracion: %v", err)
	}

	// Validar configuraciones
	if cfg.ServerPort == "" {
		return nil, fmt.Errorf("server_port es requerido")
	}
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("configuracion de la base de datos incompletas")
	}
	// Sobreescribir con variables de entorno
	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.ServerPort = port
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		cfg.DBHost = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		cfg.DBPort = dbPort
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		cfg.DBUser = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		cfg.DBPassword = dbPassword
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		cfg.DBName = dbName
	}

	return &cfg, nil
}
