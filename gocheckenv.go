package gocheckenv

import (
	"os"
)

// CheckEnv looks for required environment variables and returns a slice of strings that includes all missing variables. Len 0 slice means
// no variables are missing.
func CheckEnv(vars []string) []string {
	missing := []string{}

	for index := range vars {
		if os.Getenv(vars[index]) == "" {
			missing = append(missing, vars[index])
		}
	}

	return missing
}
