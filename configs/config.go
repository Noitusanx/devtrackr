package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	Server ServerConfig
	Database DatabaseConfig
	JWT JWTConfig
}

type ServerConfig struct {
	Port string
	Env string
}

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env: getEnv("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host: getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "devtrackr"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		},
	}
}

func getEnv(key, defaultValue string) string{
	if value:= os.Getenv(key); value != ""{
		return value;
	}
	
	return defaultValue;
}