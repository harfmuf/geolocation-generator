package path_supplier

import (
	"encoding/json"
	"errors"
	"fmt"
	m "github.com/harfmuf/geolocation-generator/model"
	"io/ioutil"
	"net/http"
)

// precision to around .0001 deg, which corresponds to max 10m
const URL_FORMAT = "%s%s?point=%.5f,%.5f&point=%.5f,%.5f&vehicle=%s&key=%s&points_encoded=false"
const GRAPHHOPPER_API_URL_BASE = "https://graphhopper.com/api/1"
const ROUTING_API_URL = "/route"

type GraphHopperPathSuplier struct {
	baseUrl string
	apiKey  string
}

type GraphHopperResponse struct {
	Paths []struct {
		Points struct {
			Coordinates [][]float64
		}
	}
}

func (g *GraphHopperPathSuplier) FindPath(from *m.Location, to *m.Location, vehicle string) ([][]*m.Location, error) {
	url := g.GetUrlFromLocations(from, to, vehicle)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request return status %d, body: %s", res.StatusCode, res.Body))
	}
	return decodeResponseToPath(res)
}

func decodeResponseToPath(response *http.Response) ([][]*m.Location, error) {
	var bodyDecoded GraphHopperResponse
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &bodyDecoded)
	if err != nil {
		return nil, err
	}
	return getPathsArrayFromJson(&bodyDecoded), nil
}

func getPathsArrayFromJson(json *GraphHopperResponse) [][]*m.Location {
	paths := make([][]*m.Location, len(json.Paths))
	for i, path := range json.Paths {
		currResult := make([]*m.Location, len(path.Points.Coordinates))
		for j, point := range path.Points.Coordinates {
			currResult[j] = &m.Location{Latitude: point[0], Longitude: point[1]}
		}
		paths[i] = currResult
	}
	return paths
}

func validateErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (g *GraphHopperPathSuplier) GetUrlFromLocations(from *m.Location, to *m.Location, vehicle string) string {
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
