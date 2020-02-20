package gocheckenv

import (
	"os"
)

// CheckEnv looks for required environment variables and exits if any are missing
func CheckEnv(vars []string) []string {
	missing := []string{}

	for index := range vars {
		if os.Getenv(vars[index]) == "" {
			missing = append(missing, vars[index])
		}
	}

	return missing
}
