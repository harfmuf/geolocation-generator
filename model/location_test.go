package model

import (
	"testing"
)

func TestDistance(t *testing.T) {
	l := Location{Latitude: 20.0, Longitude: 50.0}
	d := l.Distance(&Location{Latitude: 20.0, Longitude: 50.1})
	t.Logf("Calculated distance: %f", d)
	if d < 5475 && d > 5476 {
		t.Errorf("Wrong distance!")
	}
}
