package auth

import (
	"cli/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func New(conf *config.CliConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate to remote service.",
		RunE: func(*cobra.Command, []string) error {
			return run(conf)
		},
	}

	return cmd
}

func run(conf *config.CliConfig) error {
	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/bearer", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+conf.GetString("auth_key"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		fmt.Println("authentication failed")
	} else {
		fmt.Println("authentication succeeded")
	}

	return nil
}
