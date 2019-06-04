package conf

import (
	rand "math/rand"
	t "time"

	m "github.com/harfmuf/geolocation-generator/model"
)

//Conf represents configuration parameters
type Conf struct {
	Start      t.Time       `json:"start"`
	Stop       t.Time       `json:"stop"`
	Home       m.Location   `json:"home"`
	Work       m.Location   `json:"work"`
	Activities []m.Activity `json:"activities"`
	Interval   int8         `json:"interval"`
}

func (c *Conf) RandomActivity(weekday string) *m.Activity {
	filtered := []*m.Activity{}
	chanceSum := 0
	for _, act := range c.Activities {
		if contains(act.Weekdays, weekday) {
			filtered = append(filtered, &act)
			chanceSum += act.Chance
		}
	}
	r := rand.Intn(chanceSum)
	for _, a := range filtered {
		r -= a.Chance
		if r < 0 {
			return a
		}
	}
	return filtered[0]
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
