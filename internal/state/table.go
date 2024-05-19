package state

import (
	"computerclub/internal/events"
	"fmt"
)

// Table represents a table in the computer club
type Table struct {
	isUsed    bool // Indicates if the table is currently in use
	earnings  uint // The total earnings from the table
	usageTime int  // The total usage time of the table in minutes
}

// IsUsed returns true if the table is currently in use
func (t Table) IsUsed() bool {
	return t.isUsed
}

// ChangeStatus changes the status of the table
func (t *Table) ChangeStatus(status bool) {
	t.isUsed = status
}

// Earnings returns the total earnings from the table
func (t Table) Earnings() uint {
	return t.earnings
}

// UsageTime returns the total usage time of the table as a events.TimeFormat
func (t Table) UsageTime() events.TimeFormat {
	time, _ := events.ParseTime(fmt.Sprintf("%02d:%02d", t.usageTime/60, t.usageTime%60))
	return time
}
