package logging

import (
	"fmt"
	"io"
	"os"
)

// OutputWriter defines the output writer struct
type OutputWriter struct {
	Writer io.Writer
	Level  int
}

// Writers store all the writers of the logger
var Writers []OutputWriter

// AddStdout adds stdout to the writes with a minimum log level
func AddStdout(level int) {
	Writers = append(Writers, OutputWriter{os.Stdout, level})
}

// AddFileWriter adds a file writer to the writes with a minimum log level
func AddFileWriter(level int, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("log file open error")
	}
	Writers = append(Writers, OutputWriter{f, level})
}
