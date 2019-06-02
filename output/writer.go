package output

import "time"

type Writer interface {
	Init(string) error
	WriteEntry(TimedLocationEntry) error
	writeEntryBatch([]TimedLocationEntry) error
	finalize() error
}

type TimedLocationEntry struct {
	timestamp time.Time
	entityId  string
	deviceId  string
	latitude  float64
	longitude float64
}
