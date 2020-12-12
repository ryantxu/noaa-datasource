package noaa

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

type StationMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Lat  string `json:"lat,omitempty"`
	Lon  string `json:"lon,omitempty"`
}
