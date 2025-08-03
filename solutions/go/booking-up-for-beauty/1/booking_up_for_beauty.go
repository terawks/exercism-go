package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return 12 <= t.Hour() && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return fmt.Sprintf("You have an appointment on %s", Schedule(date).Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Now()
	return time.Date(t.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
