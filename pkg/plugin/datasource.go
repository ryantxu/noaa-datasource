package plugin

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/ryantxu/noaa-datasource/pkg/models"
	"github.com/ryantxu/noaa-datasource/pkg/noaa"
)

type Datasource struct {
	Client noaa.NOAAClient
}

func GetDatasourceServeOpts() datasource.ServeOpts {
	handler := &Datasource{
		Client: noaa.NewNOAAClient(),
	}

	return datasource.ServeOpts{
		CheckHealthHandler:  handler,
		QueryDataHandler:    handler,
		CallResourceHandler: handler,
	}
}

func (ds *Datasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "TODO!!",
	}, nil
}

func (ds *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	res := backend.NewQueryDataResponse()
	for idx := range req.Queries {
		v := &req.Queries[idx]
		q, err := models.GetNOAAQuery(v)
		if err != nil {
			res.Responses[v.RefID] = backend.DataResponse{
				Error: err,
			}
		} else {
			framer := ds.getFramer(ctx, q)
			frames, err := framer.Frames()
			res.Responses[v.RefID] = backend.DataResponse{
				Frames: frames,
				Error:  err,
			}
		}
	}
	return res, nil
}

func (ds *Datasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	return sender.Send(&backend.CallResourceResponse{
		Status: http.StatusNotFound,
		Body:   []byte("Not found"),
	})
}

func (ds *Datasource) getFramer(ctx context.Context, query *models.NOAAQuery) data.Framer {
	if query.Product == "high_low" && query.Date == "latest" {
		query.Product = "predictions"
		query.Date = "query"
		query.TimeRange.From = time.Now().UTC().Truncate(time.Minute)
		query.TimeRange.To = query.TimeRange.From.Add(time.Hour * 12)

		points, ok := ds.Client.Query(ctx, query).(*noaa.DataResponse)
		if !ok {
			return &models.ErrorFramer{Error: fmt.Errorf("erro converting")}
		}

		idx := 0
		v0, _ := strconv.ParseFloat(points.Predictions[0].V, 64)
		v1, _ := strconv.ParseFloat(points.Predictions[1].V, 64)
		ty := "Low"
		if v1 > v0 {
			ty = "High"
		}
		for i := 1; i < len(points.Predictions); i++ {
			v1, err := strconv.ParseFloat(points.Predictions[i].V, 64)
			if err == nil {
				if ty == "High" {
					if v1 < v0 {
						idx = i
						break
					}
				} else {
					if v1 > v0 {
						idx = i
						break
					}
				}
				v0 = v1
			}
		}

		row := points.Predictions[idx]
		row.TY = ty
		points.Data = []noaa.Datum{row}
		points.Predictions = []noaa.Datum{}
		query.Product = "next_high_low"

		// for _, row := range points.Data {
		// 	t, err := time.ParseInLocation("2006-01-02 15:04", row.T, time.UTC)
		// 	if err != nil && t.After(time.Now()) {
		// 		points.Data =
		// 		return points // Only the next one
		// 	}
		// }
		return points
	}
	return ds.Client.Query(ctx, query)
}
