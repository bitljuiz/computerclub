package events

import (
	"fmt"
	"time"
)

// TimeFormat represents a structure in which time is stored, in the format XX:XX.
type TimeFormat struct {
	hour    int // contains hour value (0 <= hour <= 23)
	minutes int // contains minutes value (0 <= minutes <= 59)
}

func (tf TimeFormat) String() string {
	return fmt.Sprintf("%02d:%02d", tf.hour, tf.minutes)
}

// Hour returns the hour component of the TimeFormat.
func (tf TimeFormat) Hour() int {
	return tf.hour
}

// Minutes returns the minutes component of the TimeFormat.
func (tf TimeFormat) Minutes() int {
	return tf.minutes
}

// NewTimeFormat creates a new TimeFormat instance from a standard Go time.Time object.
func NewTimeFormat(standardTimeFormat time.Time) TimeFormat {
	return TimeFormat{hour: standardTimeFormat.Hour(), minutes: standardTimeFormat.Minute()}
}

// EarlierThan compares the current TimeFormat instance with another TimeFormat instance.
// It returns true if the current instance is earlier than the other instance.
func (tf TimeFormat) EarlierThan(other TimeFormat) bool {
	return tf.hour < other.hour || (tf.hour == other.hour && tf.minutes < other.minutes)
}

// InMinutes returns the total number of minutes represented by the TimeFormat instance.
func (tf TimeFormat) InMinutes() int {
	return tf.hour*60 + tf.minutes
}

// Difference calculates the difference in minutes between the current TimeFormat instance and another TimeFormat instance.
// It returns an error if the other TimeFormat instance is earlier than the current instance.
func (tf TimeFormat) Difference(other TimeFormat) (int, error) {
	if other.EarlierThan(tf) {
		return -1, fmt.Errorf("cannot get the difference betewen, cause %v is later than %v", other, tf)
	}
	return other.InMinutes() - tf.InMinutes(), nil
}
