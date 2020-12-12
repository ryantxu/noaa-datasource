package test

import (
	"path"
	"testing"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/ryantxu/noaa-datasource/pkg/models"
	"github.com/ryantxu/noaa-datasource/pkg/testdata"
)

func TestWaterLevels(t *testing.T) {
	doWaterLevels(t).run(t)
}

var doWaterLevels testServerScenarioFn = func(t *testing.T) *testScenario {
	testFilePath := "water_level"
	return &testScenario{
		name:             testFilePath,
		mockResponsePath: path.Join("../../testdata", testFilePath+".json"),
		queries: []backend.DataQuery{
			{
				QueryType:     models.QueryTypeTidesAndCurrents,
				RefID:         "A",
				MaxDataPoints: 100,
				Interval:      1000,
				TimeRange: backend.TimeRange{
					From: time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2020, 10, 2, 0, 0, 0, 0, time.UTC),
				},
				JSON: testdata.SerializeStruct(t, &models.NOAAQuery{
					Station: 9414750,
					Product: testFilePath,
				}),
			},
		},
		goldenFileName: testFilePath,
	}
}
