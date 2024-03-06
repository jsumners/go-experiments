package conf

import (
	"cli/internal/config"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func New(cfg *config.CliConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration.",
	}

	dumpConfigCommand := &cobra.Command{
		Use:   "dump",
		Short: "Write found configuration to stdout.",
		Long: heredoc.Doc(`
			Write the configuration, as the application has read it
			from the configuration file, to stdout.
		`),
		RunE: func(*cobra.Command, []string) error {
			return dumpConfig(cfg)
		},
	}
	cmd.AddCommand(dumpConfigCommand)

	generateConfigCommand := &cobra.Command{
		Use:   "generate",
		Short: "Write default configuration to stdout.",
		RunE: func(*cobra.Command, []string) error {
			return generateConfig(cfg)
		},
	}
	cmd.AddCommand(generateConfigCommand)

	return cmd
}

func dumpConfig(cfg *config.CliConfig) error {
	yml, err := cfg.GenerateCurrentYaml()
	if err != nil {
		return err
	}
	fmt.Println(yml)
	return nil
}

func generateConfig(cfg *config.CliConfig) error {
	yml, err := cfg.GenerateDefaultYaml()
	if err != nil {
		return err
	}
	fmt.Println(yml)
	return nil
}
