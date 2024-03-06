package config

import (
	"errors"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type CliConfig struct {
	*viper.Viper `mapstructure:"-"`
	AuthKey      string `mapstructure:"auth_key"`
}

func New() *CliConfig {
	return &CliConfig{
		Viper: viper.New(),
	}
}

func (c *CliConfig) ReadConfig(configFilePath string) error {
	c.SetConfigName("conf")
	c.SetConfigType("yaml")
	c.AddConfigPath(".")
	c.AddConfigPath("/etc/cli/")
	c.AddConfigPath("$HOME/.cli/")
	c.SetEnvPrefix("CLI")
	c.AutomaticEnv()

	envFile := c.GetString("config_file")
	if configFilePath != "" {
		// This path means that `--conf-file <file>` flag has been supplied.
		// Therefore, we want to prefer it over the environment.
		c.SetConfigFile(configFilePath)
	} else if envFile != "" {
		// Fallback to the file specified by `CLI_CONFIG_FILE`.
		c.SetConfigFile(envFile)
	}

	err := c.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil
		}
		return fmt.Errorf("unable to read configuration file: %w", err)
	}

	err = c.Unmarshal(c)
	if err != nil {
		return fmt.Errorf("unable to unmarshal configuration: %w", err)
	}

	return nil
}

func (c *CliConfig) GenerateCurrentYaml() (string, error) {
	encodedData := make(map[string]any)
	err := mapstructure.Decode(c, &encodedData)
	if err != nil {
		return "", fmt.Errorf("unable to encode configuration: %w", err)
	}

	yamlEncoded, err := yaml.Marshal(encodedData)
	if err != nil {
		return "", fmt.Errorf("unable to marshal configuration to yaml: %w", err)
	}

	return string(yamlEncoded), nil
}

func (c *CliConfig) GenerateDefaultYaml() (string, error) {
	defaultConfig := CliConfig{}
	err := defaults.Set(&defaultConfig)
	if err != nil {
		return "", fmt.Errorf("unable to generate default config: %w", err)
	}

	// We decode to a generic interface in order to rename the struct fields
	// according to the `mapstructure` tag.
	var encoded map[string]any
	err = mapstructure.Decode(defaultConfig, &encoded)
	if err != nil {
		return "", fmt.Errorf("unable to decode configuration: %w", err)
	}

	yamlEncoded, err := yaml.Marshal(encoded)
	if err != nil {
		return "", fmt.Errorf("unable to encode configuration to yaml: %w", err)
	}

	return string(yamlEncoded), nil
}
