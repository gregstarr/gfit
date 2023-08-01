// Active Zone Minutes Time Series
// The Active Zone Minutes (AZM) Time Series endpoints are used for querying the user's heart-pumping activity throughout the day.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/active-zone-minutes-timeseries/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetAzmtimeSeriesByDate
// Returns the daily summary values over a period of time by specifying a date and time period.
func GetAzmtimeSeriesByDate(client *http.Client, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}

// GetAzmtimeSeriesByInterval
// Returns the daily summary values over an interval by specifying a date range.
func GetAzmtimeSeriesByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
