package utils

import "os"

func GetEnv(key, callback string) string {
	value := os.Getenv(key)
	if value == "" {
		return callback
	}
	return value
}
