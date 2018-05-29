package env

import (
	"fmt"
	"os"
)

// GetEnvVar gets an environment variable and optionally throws if not found
func GetEnvVar(key string, throw bool) string {
	value := os.Getenv(key)

	if len(value) <= 0 && throw {
		panic(fmt.Sprintf("Missing env. variable: %s", key))
	}

	return value
}
