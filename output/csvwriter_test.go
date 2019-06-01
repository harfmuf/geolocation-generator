package output

import "testing"

func TestShouldRun(t *testing.T) {
	c := CsvWriter{}
	c.init("bla")
}

func TestShouldRunWriteHeaders(t *testing.T) {
	c := CsvWriter{}
	c.writeHeaders([]string{"bla"})
}
