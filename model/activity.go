package model

type Activity struct {
	Location  Location
	Duration  float64
	Chance    int
	RouteTo   []Location
	RouteBack []Location
	Weekdays  []string
}
