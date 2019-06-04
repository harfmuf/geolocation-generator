package model

import (
	"reflect"
)

type Activity struct {
	Location  Location `json:"location"`
	Duration  float64  `json:"duration"`
	Chance    int      `json:"chance"`
	RouteTo   []*Location
	RouteBack []*Location
	Weekdays  []string `json:"weekdays"`
}

func EqualsActivity(a, b *Activity) bool {
	return EqualsLocation(&a.Location, &b.Location) &&
		a.Duration == b.Duration &&
		a.Chance == b.Chance &&
		reflect.DeepEqual(a.Weekdays, b.Weekdays)
}
