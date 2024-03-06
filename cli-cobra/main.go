package main

import (
	"cli/internal/cmd/auth"
	"cli/internal/cmd/conf"
	"cli/internal/cmd/root"
	"cli/internal/config"
	"fmt"
	"os"
)

var configFilePath string
var cfg *config.CliConfig

func main() {
	cfg = config.New()
	cmd := root.New(&configFilePath, initConfig)
	cmd.AddCommand(conf.New(cfg))
	cmd.AddCommand(auth.New(cfg))

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		// This should be the only location where we invoke os.Exit.
		os.Exit(1)
	}
}

func initConfig() error {
	return cfg.ReadConfig(configFilePath)
}
