package noaa

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/experimental"
	"github.com/ryantxu/noaa-datasource/pkg/models"
)

// view-source:https://tidesandcurrents.noaa.gov/cgi-bin/stationtideinfo.cgi?Stationid=9414750&timezone=GMT&units=english&clock=24hour&decimalPlaces=2

// 00:18 |-0.95|low
// 06:56 |5.57|high
// 11:55 |2.21|low
// 18:06 |7.76|high
// 01:07 |low|NL

type StationData struct {
	Query *models.NOAAQuery
	URL   string

	Data []Datum `json:"data,omitempty"`
}

type CGIClient struct {
	Client experimental.Client // https://tidesandcurrents.noaa.gov/cgi-bin/stationtideinfo.cgi
}

func NewCGIClient() CGIClient {
	return CGIClient{
		Client: experimental.NewRestClient("https://tidesandcurrents.noaa.gov/cgi-bin/", map[string]string{}),
	}
}

// const layoutXXX = "" // "yyyyMMdd HH:mm"

func (c *CGIClient) getQueryString(q *models.NOAAQuery) (string, error) {
	qstr := "timezone=GMT&units=english&clock=24hour&decimalPlaces=2"
	if q.Station < 100 {
		return "", fmt.Errorf("missing station")
	}
	if q.Units == "english" {
		qstr += "&units=english"
	} else {
		qstr += "&units=metric"
	}

	qstr += fmt.Sprintf("&Stationid=%d", q.Station)
	return qstr, nil
}

func (c *CGIClient) Fetch(ctx context.Context, q *models.NOAAQuery) ([]byte, error) {
	qstr, err := c.getQueryString(q)
	if err != nil {
		return nil, err
	}
	return c.Client.Fetch(ctx, "stationtideinfo.cgi", qstr)
}

func (c *CGIClient) Query(ctx context.Context, q *models.NOAAQuery) data.Framer {
	qstr, err := c.getQueryString(q)
	if err != nil {
		return &models.ErrorFramer{Error: err}
	}
	b, err := c.Client.Fetch(ctx, "stationtideinfo.cgi", qstr)
	if err != nil {
		return &models.ErrorFramer{Error: err}
	}

	val := &StationData{
		Query: q,
		URL:   qstr,
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "|")
		_, err := strconv.ParseFloat(row[1], 64)
		if err == nil && len(row) == 3 {
			val.Data = append(val.Data, Datum{
				T:  strings.TrimSpace(row[0]),
				V:  row[1],
				TY: row[2],
			})
		}
	}

	return val
}

func (s *StationData) Frames() (data.Frames, error) {
	count := len(s.Data)

	u_length := "lengthm"
	if s.Query.Units == "english" {
		u_length = "lengthft"
	}

	var fields []fieldInfo

	// Everything has T
	fields = append(fields, fieldInfo{
		name:  data.TimeSeriesTimeFieldName,
		field: data.NewFieldFromFieldType(data.FieldTypeTime, count),
		getter: func(row Datum) (interface{}, error) {
			d := time.Now().UTC().Format("2006/01/02")
			s := d + " " + row.T
			return time.ParseInLocation("2006/01/02 15:04", s, time.UTC) // Parse string
		},
	})
	fields = append(fields, fieldInfo{
		name: data.TimeSeriesValueFieldName,
		field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count).SetConfig(&data.FieldConfig{
			Unit: u_length,
		}),
		getter: func(row Datum) (interface{}, error) {
			return strconv.ParseFloat(row.V, 64) // Parse string
		},
	})
	fields = append(fields, fieldInfo{
		name:  "HighLow",
		field: data.NewFieldFromFieldType(data.FieldTypeString, count),
		getter: func(row Datum) (interface{}, error) {
			return row.TY, nil
		},
	})

	frame := data.NewFrame("")
	for i := range fields {
		fields[i].field.Name = fields[i].name
		frame.Fields = append(frame.Fields, fields[i].field)
	}

	for r := 0; r < count; r++ {
		for i := range fields {
			v, err := fields[i].getter(s.Data[r])
			if err != nil {
				continue // ???
			}
			fields[i].field.Set(r, v)
		}
	}

	frame.Name = s.Query.Product
	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: s.URL,
	}
	return data.Frames{frame}, nil
}
