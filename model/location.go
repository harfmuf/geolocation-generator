package model

import (
	"math"
	rand "math/rand"
)

const earthRadius = 6371
const oneDegree = earthRadius * 2 * math.Pi / 360 * 1000

type Location struct {
	latitude     float64
	longitude    float64
	randomRadius float64
}

func (l1 *Location) Distance(l2 *Location) float64 {
	latDist := toRadians(l2.latitude - l1.latitude)
	lonDist := toRadians(l2.longitude - l1.longitude)
	a := math.Pow(math.Sin(latDist/2), 2) +
		math.Cos(toRadians(l1.latitude))*math.Cos(toRadians(l2.latitude))*
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
		latitude: l1.latitude + r*math.Cos(th)/oneDegree,
		longitude: l1.longitude + r*math.Cos(th)/
			(oneDegree*math.Cos(l1.latitude*math.Pi/180)),
	}
}

func toRadians(deg float64) float64 {
	return deg / 360 * math.Pi
}
