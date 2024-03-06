package main

import (
	"cli/internal/client"
	"cli/internal/cmd/auth"
	"cli/internal/cmd/conf"
	"cli/internal/cmd/root"
	"cli/internal/config"
	"fmt"
	"os"
)

var configFilePath string
var cfg *config.CliConfig
var httpClient *client.Client

func main() {
	cfg = config.New()
	cmd := root.New(&configFilePath, initConfig, createClient)
	cmd.AddCommand(conf.New(cfg))
	cmd.AddCommand(auth.New(httpClient))

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

func createClient() error {
	c, err := client.New(cfg.GetString("auth_key"))
	if err != nil {
		return err
	}
	httpClient = c
	return nil
}
