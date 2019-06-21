package output

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const headers = "timestamp,latitude,longitude"
const testFile = "test.file"

func TestShouldInitAndWriteHeaders(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(testFile)
	if err != nil {
		t.Errorf("Error initializing test file")
	}
	defer c.Finalize()
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Error reading test file")
	}
	expected := []byte(headers)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect headers: %s\nexpected: %s", content, expected)
	}
}

func TestShouldWriteSingleLine(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(testFile)
	defer func() {
		err := c.Finalize()
		if err != nil {
			t.Errorf("Error closing test file\n%s", err)
		}
	}()
	if err != nil {
		t.Errorf("Error initializing test file")
	}
	now := time.Now()
	testEntry := TimedLocationEntry{
		timestamp: now,
		latitude:  50.0,
		longitude: 49.0,
	}
	line := fmt.Sprintf("\n%s,50.000000,49.000000", now)
	err = c.WriteEntry(&testEntry)
	if err != nil {
		t.Errorf("Error writing entry to test file")
	}

	content, _ := ioutil.ReadFile(testFile)
	expected := []byte(headers + line)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect values: %s\nexpected: %s", content, expected)
	}
}
func TestShouldWriteTwoLines(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(testFile)
	defer func() {
		err := c.Finalize()
		if err != nil {
			t.Errorf("Error closing test file\n%s", err)
		}
	}()
	if err != nil {
		t.Errorf("Error initializing test file")
	}
	now := time.Now()
	firstEntry := TimedLocationEntry{
		timestamp: now,
		latitude:  50.0,
		longitude: 49.0,
	}

	secondEntry := TimedLocationEntry{
		timestamp: now,
		latitude:  51.0,
		longitude: 48.0,
	}
	line1 := fmt.Sprintf("\n%s,50.000000,49.000000", now)
	line2 := fmt.Sprintf("\n%s,51.000000,48.000000", now)

	err = c.WriteEntryBatch([]*TimedLocationEntry{&firstEntry, &secondEntry})
	if err != nil {
		t.Errorf("Error writing entry to test file")
	}

	content, _ := ioutil.ReadFile(testFile)
	expected := []byte(headers + line1 + line2)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect values: %s\nexpected: %s", content, expected)
	}
}
