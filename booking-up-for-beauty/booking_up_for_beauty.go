package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    then, _ := time.Parse("1/2/2006 15:04:05", date)
    return then
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    then, _ := time.Parse("January 2, 2006 15:04:05", date)
    return then.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    then, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return then.Hour() >= 12 && then.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    then, _ := time.Parse("1/2/2006 15:04:05", date)
    return then.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
