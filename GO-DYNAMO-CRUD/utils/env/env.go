package env

import "os"

// GetEnv This function will help us get the environment variables
func GetEnv(env, defaultValue string) string {
	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}

	return environment
}
