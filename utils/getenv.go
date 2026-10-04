package utils

import "os"

func GetEnvString(field, fallback string) string {

	val := os.Getenv("REPO_DIR")

	if val == "" {
		return fallback
	}
	return val
}
