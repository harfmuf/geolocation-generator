package main

import (
	"fmt"

	m "github.com/harfmuf/geolocation-generator/model"
)

func main() {
	l := m.Location{Latitude: 20.0, Longitude: 50.0}
	fmt.Printf("l=%f;%f", l.Latitude, l.Longitude)
}
