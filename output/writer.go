package output

type Writer interface {
	init(string) error
	writeHeaders(string) error
	writeData(string) error
	finalize() error
}
