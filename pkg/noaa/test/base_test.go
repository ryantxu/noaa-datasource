package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/experimental"
	"github.com/ryantxu/noaa-datasource/pkg/models"
	"github.com/ryantxu/noaa-datasource/pkg/noaa"
	"github.com/ryantxu/noaa-datasource/pkg/plugin"
)

type testServerScenarioFn func(t *testing.T) *testScenario

type testScenario struct {
	name             string
	mockResponsePath string
	queries          []backend.DataQuery
	goldenFileName   string
	validationFn     func(t *testing.T, dr *backend.QueryDataResponse)
}

func (ts *testScenario) run(t *testing.T) {
	runTestScenario(t, ts)
}

func runTestScenario(t *testing.T, scenario *testScenario) {
	t.Run(scenario.name, func(t *testing.T) {
		ctx := context.Background()

		req := &backend.QueryDataRequest{
			PluginContext: backend.PluginContext{},
			Queries:       scenario.queries,
		}

		if _, err := os.Stat(scenario.mockResponsePath); os.IsNotExist(err) {
			std := noaa.NewNOAAClient()
			for idx := range scenario.queries {
				nq, err := models.GetNOAAQuery(&scenario.queries[idx])
				if err != nil {
					t.Fatal(err)
				}
				b, err := std.Fetch(context.Background(), nq)
				if err != nil {
					t.Fatal(err)
				}
				err = ioutil.WriteFile(scenario.mockResponsePath, b, 0600)
				if err != nil {
					t.Fatal(err)
				}
			}
		}

		datasource := &plugin.Datasource{
			Client: noaa.NOAAClient{
				Client: NewMockClient(scenario.mockResponsePath),
			},
		}
		qdr, err := datasource.QueryData(ctx, req)

		// this should always be nil, as the error is wrapped in the QueryDataResponse
		if err != nil {
			t.Fatal(err)
		}

		if scenario.validationFn != nil {
			scenario.validationFn(t, qdr)
		}

		// write out the golden for all data responses
		for i, dr := range qdr.Responses {
			fname := fmt.Sprintf("%s-%s.golden.txt", scenario.goldenFileName, i)

			// temporary fix for golden files https://github.com/grafana/grafana-plugin-sdk-go/issues/213
			for _, fr := range dr.Frames {
				if fr.Meta != nil {
					fr.Meta.Custom = nil
				}
			}

			experimental.CheckGoldenJSONResponse(t, "../../testdata/", fname, &dr, true)
		}
	})
}
