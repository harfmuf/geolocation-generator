package path_supplier

import (
	"fmt"
	m "github.com/harfmuf/geolocation-generator/model"
	"net/http"
)

// precision to around .0001 deg, which corresponds to max 10m
const URL_FORMAT = "%s%s?point=%.4f,%.4f&point=%.4f,%.4f&vehicle=%s&key=%s"
const GRAPHHOPPER_API_URL_BASE = "https://graphhopper.com/api/1"
const ROUTING_API_URL = "/route"

type PathSupplier interface {
	FindPath(m.Location, m.Location, string) [][]m.Location
}

type GraphHopperPathSuplier struct {
	apiKey string
}

func New(apiKey string) *GraphHopperPathSuplier {
	return &GraphHopperPathSuplier{apiKey: apiKey}
}

func (g GraphHopperPathSuplier) FindPath(from m.Location, to m.Location, vehicle string) [][]m.Location {
	url := g.GetUrlFromLocations(from, to, vehicle)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(fmt.Sprintf("Request failed: %d, %s", res.StatusCode, res.Body))
	}
	return decodeResponseToPath(res)
}

func decodeResponseToPath(response *http.Response) [][]m.Location {
	return nil
}

func (g GraphHopperPathSuplier) GetUrlFromLocations(from m.Location, to m.Location, vehicle string) string {
	fromLat := from.Latitude
	fromLon := from.Longitude

	toLat := to.Latitude
	toLon := to.Longitude
	return fmt.Sprintf(
		URL_FORMAT,
		GRAPHHOPPER_API_URL_BASE,
		ROUTING_API_URL,
		fromLat,
		fromLon,
		toLat,
		toLon,
		vehicle,
		g.apiKey)
}
