package events

import (
	"fmt"
)

// ID just the marker to identify that the field deals with id of events
type ID int

// Event contains func String() that all type of event should implement
type Event interface {
	String() string
}

// IncomeEvent describes a format for the events that incomes in computer's club system
type IncomeEvent struct {
	IncomeEventTime TimeFormat //contains time when the IncomeEvent happened
	IncomeEventID   ID         //contains ID of event
	IncomeEventBody string     //contains information about the event that which will need to be processed
}

func (ie IncomeEvent) String() string {
	return fmt.Sprintf("%v %d %s", ie.IncomeEventTime, ie.IncomeEventID, ie.IncomeEventBody)
}

// OutcomeEvent describes a format for the events that happened after some IncomeEvent
type OutcomeEvent struct {
	OutcomeEventTime TimeFormat //contains time when the OutcomeEvent happened
	OutcomeEventID   ID         //contains ID of event
	OutcomeEventBody string     //contains data about event that is guaranteed to be processed
}

func (oe OutcomeEvent) String() string {
	return fmt.Sprintf("%v %d %s", oe.OutcomeEventTime, oe.OutcomeEventID, oe.OutcomeEventBody)
}
