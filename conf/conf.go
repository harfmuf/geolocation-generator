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
	Work       m.Activity   `json:"work"`
	Activities []m.Activity `json:"activities"`
	Interval   int8         `json:"interval"`
}

func (c *Conf) RandomActivity(weekday string) *m.Activity {
	filtered := []*m.Activity{}
	chanceSum := 0
	for i := range c.Activities {
		if contains(c.Activities[i].Weekdays, weekday) {
			filtered = append(filtered, &c.Activities[i])
			chanceSum += c.Activities[i].Chance
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
