package output

import "time"

type Writer interface {
	Init(string) error
	WriteEntry(TimedLocationEntry) error
	WriteEntryBatch([]TimedLocationEntry) error
	Finalize() error
}

type TimedLocationEntry struct {
	timestamp time.Time
	latitude  float64
	longitude float64
}
