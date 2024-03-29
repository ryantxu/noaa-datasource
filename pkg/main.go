package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/ryantxu/noaa-datasource/pkg/plugin"
)

func main() {
	if err := datasource.Manage("ryantxu-noaa-datasource", plugin.NewNoaaInstance, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
