package test

import (
	"path"
	"testing"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/ryantxu/noaa-datasource/pkg/models"
	"github.com/ryantxu/noaa-datasource/pkg/testdata"
)

func TestHighLow(t *testing.T) {
	doHighLow(t).run(t)
}

var doHighLow testServerScenarioFn = func(t *testing.T) *testScenario {
	testFilePath := "high_low"
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
					From: time.Date(2020, 10, 1, 0, 1, 22, 0, time.UTC),
					To:   time.Date(2020, 10, 2, 0, 2, 22, 0, time.UTC),
				},
				JSON: testdata.SerializeStruct(t, &models.NOAAQuery{
					Station: 9414750,
					Product: testFilePath,
					//	Date:    "latest", //"recent", //"latest",
				}),
			},
		},
		goldenFileName: testFilePath,
	}
}
