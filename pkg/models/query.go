package models

import (
	"encoding/json"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

const (
	QueryTypeTidesAndCurrents = "TidesAndCurrents"
)

type NOAAQuery struct {
	Station string `json:"station,omitempty"`
	Product string `json:"product,omitempty"`
	Units   string `json:"units,omitempty"` //  'metric' | 'english';
	Date    string `json:"date,omitempty"`  //  'query' | 'today' | 'recent';...

	// These are added from the base query
	Interval      time.Duration     `json:"-"`
	TimeRange     backend.TimeRange `json:"-"`
	MaxDataPoints int64             `json:"-"`
	QueryType     string            `json:"-"`
}

func GetNOAAQuery(dq *backend.DataQuery) (*NOAAQuery, error) {
	query := &NOAAQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// add on the DataQuery params
	query.TimeRange = dq.TimeRange
	query.Interval = dq.Interval
	query.MaxDataPoints = dq.MaxDataPoints
	query.QueryType = dq.QueryType

	return query, nil
}
