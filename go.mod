module github.com/ryantxu/noaa-datasource

go 1.16

replace github.com/grafana/grafana-plugin-sdk-go => ../../more/grafana-plugin-sdk-go

require (
	github.com/grafana/grafana-plugin-sdk-go v0.91.1-0.20210406033415-b9e02c9c8dad
	github.com/magefile/mage v1.11.0
	github.com/stretchr/testify v1.7.0
)
