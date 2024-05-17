package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func initConfig() Config {

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("PUBLIC_HOST", "http://localhost"),
		DBPassword: getEnv("PUBLIC_HOST", "http://localhost"),
		DBAddress:  getEnv("PUBLIC_HOST", "http://localhost"),
		DBName:     getEnv("PUBLIC_HOST", "http://localhost"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {

		return value
	}

	return fallback
}