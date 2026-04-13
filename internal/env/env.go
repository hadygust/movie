package env

import "os"

func GetString(key string, fallback string) string {
	res := os.Getenv(key)

	if res == "" {
		return fallback
	}

	return res
}
