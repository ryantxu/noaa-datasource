package plugin

import (
	"context"
	"net/http"

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
	return ds.Client.Query(ctx, query)
}
