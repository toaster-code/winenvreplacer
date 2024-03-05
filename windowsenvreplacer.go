// windowsenvreplacer.go
package windowsenvreplacer

import (
	"os"
	"regexp"
)

// WindowsEnvReplacer replaces Windows environment variable references in a string.
func WindowsEnvReplacer(input string) string {
	// Define a regular expression to match expressions like %VAR_NAME%
	regex := regexp.MustCompile(`%([^%]+)%`)

	// Replace matches with their corresponding environment variable values
	result := regex.ReplaceAllStringFunc(input, func(match string) string {
		// Extract the environment variable name from the match
		varName := match[1 : len(match)-1]

		// Retrieve the Windows environment variable value
		varValue, exists := os.LookupEnv(varName)
		if !exists {
			// If the environment variable doesn't exist, keep the original match
			return match
		}

		// Replace the match with the environment variable value
		return varValue
	})

	return result
}
