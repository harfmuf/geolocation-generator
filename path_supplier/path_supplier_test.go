package path_supplier

import (
	"github.com/harfmuf/geolocation-generator/model"
	"gotest.tools/assert"
	"testing"
)

func TestUrlComposedCorrectly(t *testing.T) {
	//key := os.Getenv("graphhopperKey")
	g := New("fakeKey")
	from := model.Location{Latitude: 1.0, Longitude: 2.0}
	to := model.Location{Latitude: 6.66, Longitude: 6.66}
	vehicle := "car"
	url := g.GetUrlFromLocations(from, to, vehicle)

	assert.Equal(t, url, "https://graphhopper.com/api/1/route?point=1.0000,2.0000&point=6.6600,6.6600&vehicle=car&key=fakeKey")
}
