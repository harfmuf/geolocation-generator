package output

import (
	"fmt"
	"io"
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
	err = writeHeaders(c.file)
	if err != nil {
		return err
	}
	return nil
}

func writeHeaders(file io.Writer) error {
	t := TimedLocationEntry{}
	noHeaders := reflect.TypeOf(t).NumField()
	fieldNames := make([]string, noHeaders)
	for i := 0; i < noHeaders; i++ {
		fieldNames[i] = reflect.TypeOf(t).Field(i).Name
	}
	_, err := file.Write([]byte(strings.Join(fieldNames, ",")))
	if err != nil {
		return err
	}
	return nil
}

func (c *CsvWriter) WriteEntry(entry TimedLocationEntry) error {
	return writeInternal(c.file, entry)
}
func (c *CsvWriter) writeEntryBatch(entries []TimedLocationEntry) error {
	for _, entry := range entries {
		err := writeInternal(c.file, entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeInternal(file io.Writer, entry TimedLocationEntry) error {
	line := fmt.Sprintf("\n%s,%s,%s,%.4f,%.4f",
		entry.timestamp,
		entry.entityId,
		entry.deviceId,
		entry.latitude,
		entry.longitude)
	_, err := file.Write([]byte(line))
	return err
}

func (c *CsvWriter) finalize() error {
	return c.file.Close()
}
