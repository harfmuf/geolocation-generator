package model

import (
	"math"
	rand "math/rand"
)

const earthRadius = 6371
const oneDegree = earthRadius * 2 * math.Pi / 360 * 1000

type Location struct {
	Latitude     float64
	Longitude    float64
	RandomRadius float64
}

func (l1 *Location) Distance(l2 *Location) float64 {
	latDist := toRadians(l2.Latitude - l1.Latitude)
	lonDist := toRadians(l2.Longitude - l1.Longitude)
	a := math.Pow(math.Sin(latDist/2), 2) +
		math.Cos(toRadians(l1.Latitude))*math.Cos(toRadians(l2.Latitude))*
			math.Pow(math.Sin(lonDist/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c * 1000
}

func (l1 *Location) DistanceVia(l2 *Location, waypoints []*Location) float64 {
	distance := 0.0
	start := l1
	for _, wp := range waypoints {
		distance += start.Distance(wp)
		start = wp
	}
	distance += start.Distance(l2)
	return distance
}

func (l1 *Location) RandomLocationInDistance(meters float64) *Location {
	r := meters * math.Sqrt(rand.Float64())
	th := rand.Float64() * 2 * math.Pi

	return &Location{
		Latitude: l1.Latitude + r*math.Cos(th)/oneDegree,
		Longitude: l1.Longitude + r*math.Cos(th)/
			(oneDegree*math.Cos(l1.Latitude*math.Pi/180)),
	}
}

func toRadians(deg float64) float64 {
	return deg / 360 * math.Pi
}
