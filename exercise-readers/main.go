package main

import (
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	reader := strings.NewReader("A")
	n, _ := reader.Read(b)
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
