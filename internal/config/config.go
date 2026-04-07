package config

import "os"

type Config struct {
	Port          string
	GroqAPIKey    string
	PostgresUser  string
	PostgresPass  string
	PostgresDB    string
	PostgresPort  string
}

func Load() Config {
	return Config{
		Port:         getEnvOrDefault("PORT", "8080"),
		GroqAPIKey:   os.Getenv("GROQ_API_KEY"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:   os.Getenv("POSTGRES_DB"),
		PostgresPort: getEnvOrDefault("POSTGRES_PORT", "5432"),
	}
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
