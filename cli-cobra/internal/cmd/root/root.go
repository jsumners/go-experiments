package root

import (
	"github.com/spf13/cobra"
)

type InitFn func() error

func New(configFile *string, initFn InitFn) *cobra.Command {
	cmd := &cobra.Command{
		Short: "cli is an example",
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initFn()
		},
	}

	cmd.PersistentFlags().StringVarP(
		configFile,
		"conf-file",
		"c",
		"",
		"Set the file from which configuration will be loaded.",
	)

	return cmd
}
