package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/catalyzeio/catalyze/helpers"
	"github.com/catalyzeio/catalyze/models"
)

// ListVars lists all environment variables.
func ListVars(settings *models.Settings) {
	helpers.SignIn(settings)
	envVars := helpers.ListEnvVars(settings)
	for key, value := range envVars {
		fmt.Printf("%s=%s\n", key, value)
	}
}

// SetVar adds a new environment variables or updates the value of an existing
// environment variables. Any changes to environment variables will not take
// effect until the service is redeployed by pushing new code or via
// `catalyze redeploy`.
func SetVar(variables []string, settings *models.Settings) {
	helpers.SignIn(settings)
	envVars := helpers.ListEnvVars(settings)

	envVarsMap := make(map[string]string, len(variables))
	for _, envVar := range variables {
		pieces := strings.Split(envVar, "=")
		if len(pieces) != 2 {
			fmt.Printf("Invalid variable format. Expected <key>=<value> but got %s\n", envVar)
			os.Exit(1)
		}
		envVarsMap[pieces[0]] = pieces[1]
	}

	for key := range envVarsMap {
		if _, ok := envVars[key]; ok {
			helpers.UnsetEnvVar(key, settings)
		}
	}

	helpers.SetEnvVars(envVarsMap, settings)
	fmt.Println("Set.")
}

// UnsetVar deletes an environment variable. Any changes to environment variables
// will not take effect until the service is redeployed by pushing new code
// or via `catalyze redeploy`.
func UnsetVar(variable string, settings *models.Settings) {
	helpers.SignIn(settings)
	helpers.UnsetEnvVar(variable, settings)
	fmt.Println("Unset.")
}
