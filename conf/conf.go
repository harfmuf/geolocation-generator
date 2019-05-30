package conf

import (
	rand "math/rand"
	t "time"

	m "github.com/harfmuf/geolocation-generator/model"
)

type Conf struct {
	start      t.Time
	stop       t.Time
	home       *m.Location
	work       *m.Location
	activities []*m.Activity
	interval   int8
}

func (c *Conf) RandomActivity(weekday string) *m.Activity {
	filtered := []*m.Activity{}
	chanceSum := 0
	for _, act := range c.activities {
		if contains(act.Weekdays, weekday) {
			filtered = append(filtered, act)
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
