package conf

import (
	"testing"

	m "github.com/harfmuf/geolocation-generator/model"
)

func TestRandomActivity(t *testing.T) {
	const WEEKDAY = "MONDAY"
	acts := []m.Activity{
		m.Activity{Chance: 1, Weekdays: []string{WEEKDAY}},
		m.Activity{Chance: 2, Weekdays: []string{WEEKDAY}},
	}
	c := Conf{Activities: append(acts, m.Activity{Chance: 100, Weekdays: []string{"OTHER"}})}
	for i := 0; i < 10; i++ {
		a := c.RandomActivity(WEEKDAY)
		if !containsChance(acts, a) {
			t.Errorf("Wrong Activity{Chance:%d, Weekday:%s} for Weekday: %s", a.Chance, a.Weekdays[0], WEEKDAY)
		}
	}
}

func containsChance(s []m.Activity, e *m.Activity) bool {
	for i := range s {
		if s[i].Chance == e.Chance {
			return true
		}
	}
	return false
}
