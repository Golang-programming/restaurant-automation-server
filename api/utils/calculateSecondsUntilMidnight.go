package utils

import "time"

// CalculateSecondsUntilMidnight calculates the number of seconds until midnight.
func CalculateSecondsUntilMidnight() int {
	now := time.Now()
	nextMidnight := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
	return int(nextMidnight.Sub(now).Seconds())
}
