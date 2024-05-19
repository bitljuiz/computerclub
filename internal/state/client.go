package state

import (
	"computerclub/internal/events"
	"regexp"
)

const (
	Came      = 1 // Represents the "Came" status of a client
	Sit       = 2 // Represents the "Sit" status of a client
	IsWaiting = 3 // Represents the "IsWaiting" status of a client
	Left      = 4 // Represents the "Left" status of a client
)

// Client represents a client in the computer club
type Client struct {
	name       string            // The name of the client
	status     int               // The current status of the client
	tableIndex int               // The index of the table the client is using
	startTime  events.TimeFormat // Time, when client started using the table
}

// Name returns the name of the client
func (c Client) Name() string {
	return c.name
}

// Status returns the current status of the client
func (c Client) Status() int {
	return c.status
}

// SetStatus sets a new status for the client
func (c *Client) SetStatus(newStatus int) {
	c.status = newStatus
}

// TableIndex returns the index of the table the client is using
func (c Client) TableIndex() int {
	return c.tableIndex
}

// SetTable sets a new table index for the client
func (c *Client) SetTable(tableIndex int) {
	c.tableIndex = tableIndex
}

// StartTime returns who much time does the client use a computer
func (c Client) StartTime() events.TimeFormat {
	return c.startTime
}

// SetTime sets time when the client started using a computer
func (c *Client) SetTime(time events.TimeFormat) {
	c.startTime = time
}

// NewClient creates a new Client
func NewClient(name string, status int, tableIndex int, startTime events.TimeFormat) Client {
	return Client{name: name, status: status, tableIndex: tableIndex, startTime: startTime}
}

// NameIsValid validates the client's name
func NameIsValid(word string) bool {
	regex := `^[a-z0-9_-]+$`
	re := regexp.MustCompile(regex)
	return re.MatchString(word)
}
