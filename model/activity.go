package model

type Activity struct {
	Location  Location `json:"location"`
	Duration  float64  `json:"duration"`
	Chance    int      `json:"chance"`
	RouteTo   []Location
	RouteBack []Location
	Weekdays  []string `json:"dayTypes"`
}
