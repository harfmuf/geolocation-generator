package output

import "time"

type Writer interface {
	init(string) error
	writeEntry(TimedLocationEntry) error
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
