package app

import (
	"cli/internal/client"
	"cli/internal/config"
)

type CliApp struct {
	Config     *config.CliConfig
	HttpClient *client.Client
}
