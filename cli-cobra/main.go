package main

import (
	"cli/internal/app"
	"cli/internal/client"
	"cli/internal/commands/auth"
	"cli/internal/commands/conf"
	"cli/internal/commands/root"
	"cli/internal/config"
	"fmt"
	"os"
)

var cliApp *app.CliApp
var configFilePath string

func main() {
	cliApp = &app.CliApp{
		Config: config.New(),
	}

	cmd := root.New(&configFilePath, initConfig, createClient)
	cmd.AddCommand(conf.New(cliApp))
	cmd.AddCommand(auth.New(cliApp))

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		// This should be the only location where we invoke os.Exit.
		os.Exit(1)
	}
}

func initConfig() error {
	return cliApp.Config.ReadConfig(configFilePath)
}

func createClient() error {
	httpClient, err := client.New(cliApp.Config.GetString("auth_key"))
	if err != nil {
		return err
	}
	cliApp.HttpClient = httpClient
	return nil
}
