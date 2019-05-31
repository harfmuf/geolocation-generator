package path_supplier

import (
	"bytes"
	"github.com/harfmuf/geolocation-generator/model"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const FAKE_URL = "fakeUrl"
const FAKE_KEY = "fakeKey"

func TestUrlComposedCorrectly(t *testing.T) {
	//key := os.Getenv("graphhopperKey")
	g := New(FAKE_URL, FAKE_KEY)
	from := model.Location{Latitude: 1.0, Longitude: 2.0}
	to := model.Location{Latitude: 6.66, Longitude: 6.66}
	vehicle := "car"
	url := g.GetUrlFromLocations(from, to, vehicle)

	assert.Equal(t, url, "fakeUrl/route?point=1.0000,2.0000&point=6.6600,6.6600&vehicle=car&key=fakeKey&points_encoded=false")
}

func TestResponseDecodedCorrectly(t *testing.T) {
	//key := os.Getenv("graphhopperKey")
	g := New(FAKE_URL, FAKE_KEY)
	jsonPayload, _ := ioutil.ReadFile("sample_response.json")
	payload := []byte(jsonPayload)
	r := ioutil.NopCloser(bytes.NewReader(payload))
	response := http.Response{Body: r}
	g.DecodeResponseToPath(&response)
}
