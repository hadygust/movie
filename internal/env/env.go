package env

import (
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
