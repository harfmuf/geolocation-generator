package conf

import (
	"testing"
	"time"

	m "github.com/harfmuf/geolocation-generator/model"
)

func TestLoadConf(t *testing.T) {
	conf, _ := LoadConf("params.json")
	var timeZero time.Time
	if conf.Start == timeZero {
		t.Error("Start not properly initialized")
	}
	if conf.Stop == timeZero {
		t.Error("Stop not properly initialized!")
	}
	if conf.Interval == 0 {
		t.Error("Interval not properly initialized!")
	}
	var zeroLoc m.Location
	if m.EqualsLocation(&conf.Home, &zeroLoc) {
		t.Error("Home not properly initialized!")
	}
	var zeroAct m.Activity
	if m.EqualsActivity(&conf.Work, &zeroAct) {
		t.Error("Work not properly initialized!")
	}
	if len(conf.Activities) != 3 {
		t.Error("Activities not properly initialized!")
	}
}

func TestLoadConfError(t *testing.T) {
	conf, err := LoadConf("notExisting.json")
	if conf != nil || err == nil {
		t.Error("Expected error did not occurred!")
	}
}
