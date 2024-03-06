// winenvreplacer.go
package winenvreplacer

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// replaceEnvVariables replaces Windows environment variable references in a string.
func ReplaceEnvVariables(src string) string {
	// Define a regular expression to match expressions like %VAR_NAME%
	regex := regexp.MustCompile(`%([^%]+)%`)

	// Replace matches with their corresponding environment variable values
	result := regex.ReplaceAllStringFunc(src, func(match string) string {
		// Retrieve the Windows environment variable value using the match as the variable name
		// Extract the variable name from the match
		varName := match[1 : len(match)-1]
		envVarValue, envVarExists := os.LookupEnv(varName)

		// Check if the environment variable exists, otherwise return the original match
		if !envVarExists {
			fmt.Printf("Environment variable %s does not exist\n", match)
			return ""
		}
		fmt.Printf("Replacing %s with %s\n", match, envVarValue)
		return envVarValue
	})

	return result
}

// LookupEnvWithPercent retrieves the value of the environment variable named by the key,
func LookupEnvWithPercent(key string) (string, bool) {
	// Retrieve the environment variable value using the modified key, where
	// it uses trim to remove leading and trailing percent signs from the key string.
	// Note: If there are percent signs within the string, they will be retained.
	envVarValue, envVarExists := os.LookupEnv(strings.Trim(key, "%"))
	return envVarValue, envVarExists
}

func Example() {
	// Test the function "replaceEnvVariables":
	// Example usage
	inputString := "GOPATH=%GOPATH%, PROMPT=%PROMPT%, USERPROFILE=%USERPROFILE%, Non-existent=%NON_EXISTENT%"
	result := ReplaceEnvVariables(inputString)
	fmt.Println("INPUT:")
	fmt.Println(inputString)
	fmt.Println("RESULT:")
	fmt.Println(result)

	// Test the function "LookupEnvWithPercent"
	// Example usage
	fmt.Println(LookupEnvWithPercent("%PATH%"))
}
