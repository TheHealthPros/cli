package commands

import (
	"fmt"

	"github.com/catalyzeio/catalyze/helpers"
	"github.com/catalyzeio/catalyze/models"
)

// SupportIds prints out various IDs related to the associated environment to be
// used when contacting Catalyze support at support@catalyze.io.
func SupportIds(settings *models.Settings) {
	helpers.SignIn(settings)
	fmt.Printf(`EnvironmentID:  %s
UsersID:        %s
ServiceID:      %s
`, settings.EnvironmentID, settings.UsersID, settings.ServiceID)
}
