package path_supplier

import (
	"github.com/harfmuf/geolocation-generator/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

const FAKE_URL = "fakeUrl"
const FAKE_KEY = "fakeKey"

func TestUrlComposedCorrectly(t *testing.T) {
	g := GraphHopperPathSuplier{FAKE_URL, FAKE_KEY}
	from := model.Location{Latitude: 1.0, Longitude: 2.0}
	to := model.Location{Latitude: 6.66, Longitude: 6.66}
	vehicle := "car"
	url := g.GetUrlFromLocations(&from, &to, vehicle)

	assert.Equal(t, url, "fakeUrl/route?point=1.00000,2.00000&point=6.66000,6.66000&vehicle=car&key=fakeKey&points_encoded=false")
}

func TestResponseDecodedCorrectly(t *testing.T) {
	jsonPayload, _ := ioutil.ReadFile("sample_response.json")
	payload := []byte(jsonPayload)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(payload))
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, FAKE_KEY}

	paths := g.FindPath(&model.Location{}, &model.Location{}, "irrelevant")
	assertAlmostEqual(t, paths[0][0].Latitude, 12.416611)
	assertAlmostEqual(t, paths[1][2].Longitude, 51.132054)
}

func TestPanicOnMalformedJson(t *testing.T) {
	malformed := []byte(`}`)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(malformed)
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, FAKE_KEY}
	assert.Panics(t, func() { g.FindPath(&model.Location{}, &model.Location{}, "irrelevant") })
}

func TestPanicOnResponseStatusNotOK(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte("not_funny"))
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, FAKE_KEY}
	assert.Panics(t, func() { g.FindPath(&model.Location{}, &model.Location{}, "irrelevant") })
}

func assertAlmostEqual(t *testing.T, x float64, y float64) {
	assert.True(t, math.Abs(x-y) < 0.0001)
}
