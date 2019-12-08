package what3words

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Country      string      `json:"country"`
	Square       Square      `json:"square"`
	NearestPlace string      `json:"nearestPlace"`
	Coordinates  Coordinates `json:"coordinates"`
	Words        string      `json:"words"`
	Language     string      `json:"language"`
	MapUrl       string      `json:"map"` // map にするとエラーになるため
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Square struct {
	southwest Coordinates `json:"southwest"`
	northwest Coordinates `json:"northwest"`
}

func DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
