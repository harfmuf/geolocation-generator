package model

import (
	"testing"
)

func TestDistance(t *testing.T) {
	l := Location{Latitude: 50.0, Longitude: 20.0}
	d := l.Distance(&Location{Latitude: 50.0, Longitude: 20.1})
	t.Logf("Calculated distance: %f", d)
	if d < 5038 || d > 5039 {
		t.Errorf("Wrong distance!")
	}
}

func TestDistanceVia(t *testing.T) {
	l := Location{Latitude: 50.0, Longitude: 20.0}
	waypoints := []*Location{
		&Location{Latitude: 50.0, Longitude: 20.1},
		&Location{Latitude: 50.1, Longitude: 20.1},
	}
	d := l.DistanceVia(&Location{Latitude: 50.1, Longitude: 20.0}, waypoints)
	t.Logf("Calculated distance: %f", d)
	if d < 15635 || d > 15636 {
		t.Errorf("Wrong distance!")
	}
}

func TestRandomLocationInDistance(t *testing.T) {
	cases := [][]interface{}{
		{50.0, 20.0, 20.0},
		{70.0, 0.0, 100.0},
		{0.0, -100.0, 50.0},
		{-50.0, 60.0, 100.0},
	}
	for _, test := range cases {
		l := &Location{Latitude: test[0].(float64), Longitude: test[1].(float64)}
		p := l.RandomLocationInDistance(test[2].(float64))
		d := l.Distance(p)
		if d > test[2].(float64) {
			t.Errorf("Point is not in radius!\nExpected: %f\nActual:%f", test[2].(float64), d)
		}
	}
}
