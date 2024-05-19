package events

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// timeFormat describes a format for the time that is used everywhere in project
const (
	timeFormat = "15:04"
)

// EventFactory is a struct that is used to create new IncomeEvent object's
type EventFactory struct{}

// NewEventFactory creates new EventFactory
func NewEventFactory() EventFactory {
	return EventFactory{}
}

// ParseTime processes a given string and returns TimeFormat variable in the timeFormat format
// Returns error when it cannot parse the timeString into timeFormat format
func ParseTime(timeString string) (TimeFormat, error) {
	timeNotFormatted, err := time.Parse(timeFormat, timeString)
	if err != nil {
		return TimeFormat{}, err
	}
	return NewTimeFormat(timeNotFormatted), nil
}

// CreateEvent create new IncomeEvent from the given data string
// Returns error when the given data cannot be converted into IncomeEvent format
func (f *EventFactory) CreateEvent(data string) (IncomeEvent, error) {
	var (
		event      IncomeEvent
		timeString string
		eventBody  string
	)

	tokens := strings.SplitN(data, " ", 3)
	if len(tokens) != 3 {
		return IncomeEvent{}, fmt.Errorf("Dlina ne ta....")
	}
	timeString = tokens[0]
	eventID, err := strconv.Atoi(tokens[1])
	if err != nil {
		return IncomeEvent{}, err
	}
	eventBody = tokens[2]
	event.IncomeEventTime, err = ParseTime(timeString)
	if err != nil {
		return event, err
	}
	event.IncomeEventID = ID(eventID)
	event.IncomeEventBody = eventBody

	return event, nil
}
