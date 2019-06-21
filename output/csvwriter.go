package output

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type CsvWriter struct {
	file *os.File
}

func (c *CsvWriter) Init(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	c.file = file
	err = c.writeHeaders()
	if err != nil {
		return err
	}
	return nil
}

func (c *CsvWriter) writeHeaders() error {
	t := TimedLocationEntry{}
	noHeaders := reflect.TypeOf(t).NumField()
	fieldNames := make([]string, noHeaders)
	for i := 0; i < noHeaders; i++ {
		fieldNames[i] = reflect.TypeOf(t).Field(i).Name
	}
	_, err := c.file.Write([]byte(strings.Join(fieldNames, ",")))
	if err != nil {
		return err
	}
	return nil
}

func (c *CsvWriter) WriteEntry(entry *TimedLocationEntry) error {
	return c.writeInternal(entry)
}
func (c *CsvWriter) WriteEntryBatch(entries []*TimedLocationEntry) error {
	for _, entry := range entries {
		err := c.writeInternal(entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CsvWriter) writeInternal(entry *TimedLocationEntry) error {
	const writeFormat = "\n%s,%.6f,%.6f"
	line := fmt.Sprintf(writeFormat,
		entry.timestamp,
		entry.latitude,
		entry.longitude)
	_, err := c.file.Write([]byte(line))
	return err
}

func (c *CsvWriter) Finalize() error {
	return c.file.Close()
}
