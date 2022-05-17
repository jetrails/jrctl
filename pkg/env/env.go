package env

import (
	"os"
	"strconv"
	"strings"
)

var EnvPrefix = ""

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

func GetInt(key string, fallback int) int {
	if value := os.Getenv(normalizeKey(key)); normalizeValue(value) != "" {
		if v, err := strconv.Atoi(normalizeValue(value)); err == nil {
			return v
		}
	}
	return fallback
}

func GetString(key string, fallback string) string {
	if value := os.Getenv(normalizeKey(key)); normalizeValue(value) != "" {
		return value
	}
	return fallback
}
