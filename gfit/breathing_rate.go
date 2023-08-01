// Breathing Rate
// Breathing Rate or Respiratory Rate is a measurement of your average breaths per minute at night.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/breathing-rate/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetBreathingRateSummaryByDate
// This endpoint returns average breathing rate data for a single date. Breathing Rate data applies specifically to a user’s “main sleep,” which is the longest single period of time during which they were asleep on a given date.
func GetBreathingRateSummaryByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/br/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetBreathingRateSummaryByInterval
// This endpoint returns average breathing rate data for a date range. Breathing Rate data applies specifically to a user’s “main sleep,” which is the longest single period of time during which they were asleep on a given date.
func GetBreathingRateSummaryByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/br/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
