package output

type CsvWriter struct {
}

func (c CsvWriter) init(filename string) error {
	return nil
}

func (c CsvWriter) writeHeaders(headers []string) error {
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
