package path_supplier

import (
	"encoding/json"
	"fmt"
	m "github.com/harfmuf/geolocation-generator/model"
	"io/ioutil"
	"net/http"
)

// precision to around .0001 deg, which corresponds to max 10m
const URL_FORMAT = "%s%s?point=%.4f,%.4f&point=%.4f,%.4f&vehicle=%s&key=%s&points_encoded=false"
const GRAPHHOPPER_API_URL_BASE = "https://graphhopper.com/api/1"
const ROUTING_API_URL = "/route"

type PathSupplier interface {
	FindPath(m.Location, m.Location, string) [][]m.Location
}

type GraphHopperPathSuplier struct {
	baseUrl string
	apiKey  string
}

type Points struct {
	Coordinates [][]float64
}

type Path struct {
	Points Points
}

type GraphHopperResponse struct {
	Paths []Path
}

func New(baseUrl string, apiKey string) *GraphHopperPathSuplier {
	return &GraphHopperPathSuplier{apiKey: apiKey, baseUrl: baseUrl}
}

func (g GraphHopperPathSuplier) FindPath(from m.Location, to m.Location, vehicle string) [][]m.Location {
	url := g.GetUrlFromLocations(from, to, vehicle)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(fmt.Sprintf("Request failed: %d, %s", res.StatusCode, res.Body))
	}
	return g.DecodeResponseToPath(res)
}

func (g GraphHopperPathSuplier) DecodeResponseToPath(response *http.Response) [][]m.Location {
	var bodyDecoded GraphHopperResponse
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	validateErr(err)
	err = json.Unmarshal(body, &bodyDecoded)
	result := make([][]m.Location, len(bodyDecoded.Paths))
	for i, path := range bodyDecoded.Paths {
		currResult := make([]m.Location, len(path.Points.Coordinates))
		for j, point := range path.Points.Coordinates {
			currResult[j] = m.Location{Latitude: point[0], Longitude: point[1]}
		}
		result[i] = currResult
	}
	return result
}

func validateErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (g GraphHopperPathSuplier) GetUrlFromLocations(from m.Location, to m.Location, vehicle string) string {
	fromLat := from.Latitude
	fromLon := from.Longitude

	toLat := to.Latitude
	toLon := to.Longitude
	return fmt.Sprintf(
		URL_FORMAT,
		g.baseUrl,
		ROUTING_API_URL,
		fromLat,
		fromLon,
		toLat,
		toLon,
		vehicle,
		g.apiKey)
}
