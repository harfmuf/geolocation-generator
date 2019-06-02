package output

import (
	"io"
	"os"
	"reflect"
	"strings"
)

type CsvWriter struct {
	file *os.File
}

func (c CsvWriter) Init(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	c.file = file
	WriteHeaders(c.file)
	return nil
}

func WriteHeaders(file io.Writer) error {
	t := TimedLocationEntry{}
	noHeaders := reflect.TypeOf(t).NumField()
	fieldNames := make([]string, noHeaders)
	for i := 0; i < noHeaders; i++ {
		fieldNames[i] = reflect.TypeOf(t).Field(i).Name
	}
	file.Write([]byte(strings.Join(fieldNames, ",")))
	return nil
}

func (c CsvWriter) writeEntry(TimedLocationEntry) error {
	return nil

}
func (c CsvWriter) writeEntryBatch([]TimedLocationEntry) error {
	return nil
}
func (c CsvWriter) finalize() error {
	return nil
}
