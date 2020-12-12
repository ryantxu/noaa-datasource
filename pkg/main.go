package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/ryantxu/noaa-datasource/pkg/plugin"
)

func main() {
	backend.SetupPluginEnvironment("oas-datasource")

	err := datasource.Serve(plugin.GetDatasourceServeOpts())

	// Log any error if we could start the plugin.
	if err != nil {
		backend.Logger.Error(err.Error())
		os.Exit(1)
	}
}
