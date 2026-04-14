package env

import (
	"errors"
	"fmt"
	"os"
)

func GetString(key string, fallback string) string {
	res := os.Getenv(key)
	fmt.Println(res, " get env result")

	if res == "" {
		return fallback
	}

	return res
}

func GetSecret() (string, error) {
	res := os.Getenv("JWT_SECRET")
	if res == "" {
		return res, errors.New("secret key not found")
	}

	return res, nil
}
