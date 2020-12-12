package noaa

import (
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
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

// "predictions": [
//         {
//             "t": "2020-12-11 00:42",
//             "v": "1.553"
//         },

const dateLayout = "2006-01-02 15:04"

type BaseInfo struct {
	Query string `json:"query,omitempty"`
}

type PredictionDatum struct {
	T string `json:"t,omitempty"`
	V string `json:"v,omitempty"`
}

type PredictionResponse struct {
	BaseInfo
	Data []PredictionDatum `json:"predictions,omitempty"`
}

func (r *PredictionResponse) Frames() (data.Frames, error) {
	count := len(r.Data)

	frame := data.NewFrameOfFieldTypes("", count, data.FieldTypeTime, data.FieldTypeFloat64)
	for i := 0; i < count; i++ {
		t, err := time.Parse(dateLayout, r.Data[i].T)
		if err != nil {
			return nil, err
		}
		v, err := strconv.ParseFloat(r.Data[i].V, 64)
		if err != nil {
			return nil, err
		}

		frame.Fields[0].Set(i, t)
		frame.Fields[1].Set(i, v)
	}
	frame.Fields[0].Name = data.TimeSeriesTimeFieldName
	frame.Fields[1].Name = data.TimeSeriesValueFieldName
	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: r.Query,
	}
	return data.Frames{frame}, nil
}
