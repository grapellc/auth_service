package utils

import "time"

// Now returns the current time
func Now() time.Time {
	return time.Now()
}

// UTCNow returns the current time in UTC
func UTCNow() time.Time {
	return time.Now().UTC()
}

// FormatTime formats time to a string
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}
