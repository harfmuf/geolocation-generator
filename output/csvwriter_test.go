package output

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const HEADERS = "timestamp,entityId,deviceId,latitude,longitude"
const TEST_FILE = "test.file"

func TestShouldInitAndWriteHeaders(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(TEST_FILE)
	if err != nil {
		t.Errorf("Error initializing test file")
	}
	defer c.finalize()
	content, err := ioutil.ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("Error reading test file")
	}
	expected := []byte(HEADERS)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect headers: %s\nexpected: %s", content, expected)
	}
}

func TestShouldWriteSingleLine(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(TEST_FILE)
	defer func() {
		err := c.finalize()
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
		entityId:  "testId",
		deviceId:  "testDevice",
		latitude:  50.0,
		longitude: 49.0,
	}
	line := fmt.Sprintf("\n%s,testId,testDevice,50.0000,49.0000", now)
	err = c.WriteEntry(testEntry)
	if err != nil {
		t.Errorf("Error writing entry to test file")
	}

	content, _ := ioutil.ReadFile(TEST_FILE)
	expected := []byte(HEADERS + line)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect values: %s\nexpected: %s", content, expected)
	}
}
func TestShouldWriteTwoLines(t *testing.T) {
	c := CsvWriter{}
	err := c.Init(TEST_FILE)
	defer func() {
		err := c.finalize()
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
		entityId:  "testId",
		deviceId:  "testDevice",
		latitude:  50.0,
		longitude: 49.0,
	}

	secondEntry := TimedLocationEntry{
		timestamp: now,
		entityId:  "testId1",
		deviceId:  "testDevice2",
		latitude:  51.0,
		longitude: 48.0,
	}
	line1 := fmt.Sprintf("\n%s,testId,testDevice,50.0000,49.0000", now)
	line2 := fmt.Sprintf("\n%s,testId1,testDevice2,51.0000,48.0000", now)

	err = c.writeEntryBatch([]TimedLocationEntry{firstEntry, secondEntry})
	if err != nil {
		t.Errorf("Error writing entry to test file")
	}

	content, _ := ioutil.ReadFile(TEST_FILE)
	expected := []byte(HEADERS + line1 + line2)
	if bytes.Compare(content, expected) != 0 {
		t.Errorf("Incorrect values: %s\nexpected: %s", content, expected)
	}
}
