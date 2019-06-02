package output

import (
	"bytes"
	"testing"
)

func TestShouldRun(t *testing.T) {
	//c := CsvWriter{}
	//c.Init("bla")
}

func TestShouldWriteHeaders(t *testing.T) {
	mockFile := new(bytes.Buffer)
	WriteHeaders(mockFile)
	content := mockFile.String()
	if content != "timestamp,entityId,deviceId,latitude,longitude" {
		t.Errorf("Incorrect headers: %s", content)
	}
}
