package configs

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Estructura de la configuracion

type Config struct {
	ServerPort string `yaml:"server_port"`
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

// Cargar configuraciones desde el archivo y variables de entorno
func LoadConfig() (*Config, error) {
	configPath := getConfigPath()

	// Leer configuraciones desde el archivo YAML

	cfg, err := loadConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}

	// Validar configuracion
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	// Sobrescribir con variables de entorno
	overrideWithEnv(cfg)
	return cfg, nil
}

// Obtiene la ruta del archivo de configuracion
func getConfigPath() string {
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		return path
	}
	return "configs/config.yaml"

}

// Cargar la configuracion desde el archivo YAML
func loadConfigFromFile(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("❌error al leer el archivo de configuraciones: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("❌error al decodificar el archivo de configuracion: %v", err)

	}
	return &cfg, nil
}

// Validar los valores esenciales de la configuracion
func validateConfig(cfg *Config) error {
	if cfg.ServerPort == "" {
		return fmt.Errorf("❌server_port es requerido")
	}
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return fmt.Errorf("❌configuracion de la base de datos incompletas")
	}
	return nil
}

// Sobrescribir la configuracion de variables de entorno
func overrideWithEnv(cfg *Config) {
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
	log.Println("✅ Variables de entorno aplicadas (si existen)")
}
