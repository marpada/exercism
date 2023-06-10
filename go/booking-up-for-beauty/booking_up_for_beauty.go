package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	d, error := time.Parse("January 2, 2006 15:04:05", date)
	if error != nil {
		panic(error)
	}
	return (time.Now().Compare(d) >= 0)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	d, error := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if error != nil {
		panic(error)
	}
	return (d.Hour() >= 12) && (d.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	d := Schedule(date)
	return "You have an appointment on " + d.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
