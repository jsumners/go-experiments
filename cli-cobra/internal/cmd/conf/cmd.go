package conf

import (
	"cli/internal/config"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
			return generateConfig()
		},
	}
	cmd.AddCommand(generateConfigCommand)

	return cmd
}

func dumpConfig(cfg *config.CliConfig) error {
	encodedData := make(map[string]any)
	err := mapstructure.Decode(cfg, &encodedData)
	if err != nil {
		return fmt.Errorf("unable to encode configuration: %w", err)
	}

	data, err := yaml.Marshal(encodedData)
	if err != nil {
		return fmt.Errorf("unable to marshal configuration to yaml: %w", err)
	}

	fmt.Println(string(data))
	return nil
}

func generateConfig() error {
	defaultConfig := config.CliConfig{}
	err := defaults.Set(&defaultConfig)
	if err != nil {
		return fmt.Errorf("unable to generate default config: %w", err)
	}

	// We decode to a generic interface in order to rename the struct fields
	// according to the `mapstructure` tag.
	var encoded map[string]any
	err = mapstructure.Decode(defaultConfig, &encoded)
	if err != nil {
		return fmt.Errorf("unable to decode configuration: %w", err)
	}

	yamlEncoded, err := yaml.Marshal(encoded)
	if err != nil {
		return fmt.Errorf("unable to encode configuration to yaml: %w", err)
	}
	fmt.Println(string(yamlEncoded))

	return nil
}
