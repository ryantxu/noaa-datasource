package noaa

import (
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
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
// 		"t": "2020-12-11 00:48",
// 		"v": "1.571",
// 		"s": "0.008",
// 		"f": "1,0,0,0",
// 		"q": "p"
// 	},
// 	{

type StationMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Lat  string `json:"lat,omitempty"`
	Lon  string `json:"lon,omitempty"`
}

type Datum struct {
	T  string `json:"t,omitempty"`
	V  string `json:"v,omitempty"`
	S  string `json:"s,omitempty"`
	TY string `json:"ty,omitempty"`
	D  string `json:"d,omitempty"`
	DR string `json:"dr,omitempty"`
	G  string `json:"g,omitempty"`
	F  string `json:"f,omitempty"`
	Q  string `json:"q,omitempty"`
}

// {"error": {"message":"No data was found. This product may not be offered at this station at the requested time."}}

type ResponseError struct {
	Message string `json:"message,omitempty"`
}

type DataResponse struct {
	Query *models.NOAAQuery
	URL   string

	Metadata    StationMeta   `json:"metadata,omitempty"`
	Data        []Datum       `json:"data,omitempty"`
	Predictions []Datum       `json:"predictions,omitempty"`
	Error       ResponseError `json:"error,omitempty"`
}

type fieldInfo struct {
	name   string
	field  *data.Field
	getter func(row Datum) (interface{}, error)
}

func (r *DataResponse) Frames() (data.Frames, error) {
	arr := r.Data
	count := len(r.Data)
	if len(r.Predictions) > 0 {
		count = len(r.Predictions)
		arr = r.Predictions
	}

	var fields []fieldInfo

	// Everything has T
	fields = append(fields, fieldInfo{
		name:  data.TimeSeriesTimeFieldName,
		field: data.NewFieldFromFieldType(data.FieldTypeTime, count),
		getter: func(row Datum) (interface{}, error) {
			return time.ParseInLocation("2006-01-02 15:04", row.T, time.UTC) // Parse string
		},
	})

	u_length := "lengthm"
	if r.Query.Units == "english" {
		u_length = "lengthft"
	}

	switch r.Query.Product {
	case "next_high_low":
		fallthrough
	case "high_low":
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
	case "wind":
		fields = append(fields, fieldInfo{
			name: "Speed",
			field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count).SetConfig(&data.FieldConfig{
				Unit: "knts",
			}),
			getter: func(row Datum) (interface{}, error) {
				return strconv.ParseFloat(row.S, 64) // Parse string
			},
		})
		fields = append(fields, fieldInfo{
			name: "Direction",
			field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count).SetConfig(&data.FieldConfig{
				Unit: "degree",
			}),
			getter: func(row Datum) (interface{}, error) {
				return strconv.ParseFloat(row.D, 64) // Parse string
			},
		})
		fields = append(fields, fieldInfo{
			name:  "Direction",
			field: data.NewFieldFromFieldType(data.FieldTypeString, count),
			getter: func(row Datum) (interface{}, error) {
				return row.DR, nil
			},
		})
		fields = append(fields, fieldInfo{
			name: "Gusts",
			field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count).SetConfig(&data.FieldConfig{
				Unit: "knts",
			}),
			getter: func(row Datum) (interface{}, error) {
				return strconv.ParseFloat(row.G, 64) // Parse string
			},
		})
	case "predictions":
		fields = append(fields, fieldInfo{
			name: data.TimeSeriesValueFieldName,
			field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count).SetConfig(&data.FieldConfig{
				Unit: u_length,
			}),
			getter: func(row Datum) (interface{}, error) {
				return strconv.ParseFloat(row.V, 64) // Parse string
			},
		})
	default:
		fields = append(fields, fieldInfo{
			name:  data.TimeSeriesValueFieldName,
			field: data.NewFieldFromFieldType(data.FieldTypeFloat64, count),
			getter: func(row Datum) (interface{}, error) {
				return strconv.ParseFloat(row.V, 64) // Parse string
			},
		})
	}

	frame := data.NewFrame("")
	for i := range fields {
		fields[i].field.Name = fields[i].name
		frame.Fields = append(frame.Fields, fields[i].field)
	}

	for r := 0; r < count; r++ {
		for i := range fields {
			v, err := fields[i].getter(arr[r])
			if err != nil {
				continue // ???
			}
			fields[i].field.Set(r, v)
		}
	}

	frame.Name = r.Query.Product
	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: r.URL,
		Custom:              r.Metadata,
	}

	if r.Error.Message != "" {
		frame.AppendNotices(data.Notice{
			Severity: data.NoticeSeverityError,
			Text:     r.Error.Message,
		})
	}

	return data.Frames{frame}, nil
}
