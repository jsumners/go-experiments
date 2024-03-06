package auth

import (
	"cli/internal/client"
	"fmt"
	"github.com/spf13/cobra"
)

func New(httpClient *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate to remote service.",
		RunE: func(*cobra.Command, []string) error {
			return run(httpClient)
		},
	}

	return cmd
}

func run(httpClient *client.Client) error {
	err := httpClient.Authenticate()

	if err != nil {
		fmt.Println("authentication failed")
	} else {
		fmt.Println("authentication succeeded")
	}

	return nil
}
