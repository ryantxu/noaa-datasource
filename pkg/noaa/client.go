package noaa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/experimental"
	"github.com/ryantxu/noaa-datasource/pkg/models"
)

// "metadata": {
// 	"id": "9414750",
// 	"name": "Alameda",
// 	"lat": "37.7717",
// 	"lon": "-122.3000"
// },
// "data": [
// 	{
// 		"t": "2020-12-11 00:18",
// 		"v": "1.408",
// 		"s": "0.006",
// 		"f": "0,0,0,0",
// 		"q": "p"
// 	},

// "t": "2020-12-11 00:36",
// "s": "2.30",
// "d": "208.00",
// "dr": "SSW",
// "g": "3.10",
// "f": "0,0"

type NOAAClient struct {
	Client experimental.Client // https://api.tidesandcurrents.noaa.gov/api/prod/
}

func NewNOAAClient() NOAAClient {
	return NOAAClient{
		Client: experimental.NewRestClient("https://api.tidesandcurrents.noaa.gov/api/prod/", map[string]string{}),
	}
}

const layoutXXX = "20060102 15:04" // "yyyyMMdd HH:mm"

func getQueryString(q *models.NOAAQuery) (string, error) {
	qstr := "units=metric&time_zone=gmt&application=Grafana&format=json&datum=STND"
	if q.Station < 100 {
		return "", fmt.Errorf("missing station")
	}
	if q.Product == "" {
		return "", fmt.Errorf("missing product")
	}

	qstr += fmt.Sprintf("&begin_date=%s", url.QueryEscape(q.TimeRange.From.Format(layoutXXX)))
	qstr += fmt.Sprintf("&end_date=%s", url.QueryEscape(q.TimeRange.To.Format(layoutXXX)))
	// qstr += "&range=24"

	qstr += fmt.Sprintf("&station=%d", q.Station)
	qstr += fmt.Sprintf("&product=%s", q.Product)
	return qstr, nil
}

func (c *NOAAClient) Fetch(ctx context.Context, q *models.NOAAQuery) ([]byte, error) {
	qstr, err := getQueryString(q)
	if err != nil {
		return nil, err
	}
	return c.Client.Fetch(ctx, "datagetter", qstr)
}

func (c *NOAAClient) Query(ctx context.Context, q *models.NOAAQuery) data.Framer {
	qstr, err := getQueryString(q)
	if err != nil {
		return &models.ErrorFramer{Error: err}
	}
	bytes, err := c.Client.Fetch(ctx, "datagetter", qstr)
	if err != nil {
		return &models.ErrorFramer{Error: err}
	}

	var val data.Framer
	switch q.Product {
	case "predictions":
		val = &PredictionResponse{}
	case "water_level":
		val = &WaterLevelResponse{}
	default:
		return &models.ErrorFramer{Error: fmt.Errorf("unknown product: %s", q.Product)}
	}

	if err := json.Unmarshal(bytes, val); err != nil {
		return &models.ErrorFramer{Error: err}
	}
	return val
}
