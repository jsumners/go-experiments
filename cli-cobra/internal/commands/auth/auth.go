package auth

import (
	"cli/internal/app"
	"cli/internal/client"
	"fmt"
	"github.com/spf13/cobra"
)

func New(app *app.CliApp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate to remote service.",
		RunE: func(*cobra.Command, []string) error {
			return run(app.HttpClient)
		},
	}

	return cmd
}

func run(client *client.Client) error {
	err := client.Authenticate()

	if err != nil {
		fmt.Println("authentication failed")
	} else {
		fmt.Println("authentication succeeded")
	}

	return nil
}
