package path_supplier

import (
	"github.com/harfmuf/geolocation-generator/model"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

const fakeUrl = "fakeUrl"
const fakeKey = "fakeKey"
const precisionOneTenThousandth = 0.0001

func TestUrlComposedCorrectly(t *testing.T) {
	g := GraphHopperPathSuplier{fakeUrl, fakeKey}
	from := model.Location{Latitude: 1.0, Longitude: 2.0}
	to := model.Location{Latitude: 6.66, Longitude: 6.66}
	vehicle := "car"
	url := g.GetUrlFromLocations(&from, &to, vehicle)

	expected := "fakeUrl/route?point=1.000000,2.000000&point=6.660000,6.660000&vehicle=car&key=fakeKey&points_encoded=false"
	if url != expected {
		t.Fatalf("Expected url: %s\nActual: %s", expected, url)
	}
}

func TestResponseDecodedCorrectly(t *testing.T) {
	jsonPayload, _ := ioutil.ReadFile("sample_response.json")
	payload := []byte(jsonPayload)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(payload))
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, fakeKey}

	paths, _ := g.FindPath(&model.Location{}, &model.Location{}, "irrelevant")
	assertAlmostEqual(t, paths[0][0].Latitude, 12.416611, precisionOneTenThousandth)
	assertAlmostEqual(t, paths[1][2].Longitude, 51.132054, precisionOneTenThousandth)
}

func TestErrorOnMalformedJson(t *testing.T) {
	malformed := []byte(`}`)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(malformed)
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, fakeKey}
	_, err := g.FindPath(&model.Location{}, &model.Location{}, "irrelevant")
	if err == nil {
		t.Fatalf("Should return error for malformed json. Error not returned")
	}
}

func TestErrorOnResponseStatusNotOK(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte("not_funny"))
	}))
	defer func() { testServer.Close() }()
	g := GraphHopperPathSuplier{testServer.URL, fakeKey}
	_, err := g.FindPath(&model.Location{}, &model.Location{}, "irrelevant")
	if err == nil {
		t.Fatalf("Should return error for status not 200. Error not returned")
	}
}

// assumption: for most cases precision within 0.000000001 would be sufficient
func assertAlmostEqual(t *testing.T, x float64, y float64, precision float64) {
	if math.Abs(x-y) >= precision {
		t.Fatalf("Numbers not withing %.9f", precision)
	}
}
