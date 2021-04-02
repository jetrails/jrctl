package env

import (
	"os"
	"strings"
)

var EnvPrefix = "JR_"

func normalizeKey(key string) string {
	return EnvPrefix + strings.TrimSpace(strings.ToUpper(key))
}

func normalizeValue(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func GetBool(key string, fallback bool) bool {
	if value := os.Getenv(normalizeKey(key)); normalizeValue(value) != "" {
		return normalizeValue(value) == "true"
	}
	return fallback
}

func GetString(key string, fallback string) string {
	if value := os.Getenv(normalizeKey(key)); normalizeValue(value) != "" {
		return value
	}
	return fallback
}
