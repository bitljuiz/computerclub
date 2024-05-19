package reader

import (
	"bufio"
	"computerclub/internal/events"
	"io"
	"os"
)

// EventReader is an interface for reading events from various sources
type EventReader interface {
	// ReadLine reads a single line from the source and returns it as a string
	// Returns an error if there is an issue reading the line
	ReadLine() (string, error)

	// ReadEvent reads an event from the source and returns it as an events.IncomeEvent
	// Returns an error if there is an issue reading or parsing the event
	ReadEvent() (events.IncomeEvent, error)

	// Close closes the reader and releases any resources associated with it
	Close()
}

// FileEventReader is an implementation of EventReader that reads events from a file
type FileEventReader struct {
	readFile    *os.File       // The file being read
	fileScanner *bufio.Scanner // Scanner to read the file
}

// Close closes the file associated with the FileEventReader
func (f *FileEventReader) Close() {
	f.readFile.Close()
}

// ReadLine reads the next line from the file and returns it as a string
// Returns an error if there is an issue reading the line or if the end of the file is reached (io.EOF)
func (f *FileEventReader) ReadLine() (string, error) {
	if f.fileScanner.Scan() {
		return f.fileScanner.Text(), nil
	}
	if err := f.fileScanner.Err(); err != nil {
		return "", err
	}
	return "", io.EOF
}

// ReadEvent reads the next event from the file and returns it as an events.IncomeEvent.
// Returns an error if there is an issue reading or parsing the event.
func (f *FileEventReader) ReadEvent() (events.IncomeEvent, error) {
	var event events.IncomeEvent
	line, err := f.ReadLine()
	if err != nil {
		return event, err
	}

	factory := events.NewEventFactory()
	event, err = factory.CreateEvent(line)
	if err != nil {
		return event, err
	}
	return event, nil
}

// NewFileEventReader creates a new FileEventReader for the file located by filepath string.
// Returns an error if there is an issue opening the file.
func NewFileEventReader(filepath string) (*FileEventReader, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	return &FileEventReader{readFile: readFile, fileScanner: fileScanner}, nil
}
