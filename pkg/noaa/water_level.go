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
// 		"t": "2020-12-11 00:48",
// 		"v": "1.571",
// 		"s": "0.008",
// 		"f": "1,0,0,0",
// 		"q": "p"
// 	},
// 	{

type WaterLevelDatum struct {
	T string `json:"t,omitempty"`
	V string `json:"v,omitempty"`
}

type WaterLevelResponse struct {
	BaseInfo
	Metadata StationMeta       `json:"metadata,omitempty"`
	Data     []WaterLevelDatum `json:"data,omitempty"`
}

func (r *WaterLevelResponse) Frames() (data.Frames, error) {
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
	frame.Name = r.Metadata.Name
	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: r.Query,
		Custom:              r.Metadata,
	}
	return data.Frames{frame}, nil
}
