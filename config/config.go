package config

import (
	"os"
)

func Required(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("required environment variable not set: " + key)
	}
	return v
}

func Optional(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func Port() string {
	return Optional("PORT", "8080")
}

func APIKey() string {
	return Required("API_KEY")
}
