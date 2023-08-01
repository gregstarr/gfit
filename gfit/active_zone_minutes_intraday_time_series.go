// Active Zone Minutes Intraday Time Series
// The Active Zone Minutes (AZM) Intraday Time Series endpoints retrieve the user's heart-pumping activity for a specific date or 24 hour period. Intraday support can extend the detail-level response to include 1min, 5min and 15min for Active Zone Minutes.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetAzmbyIntervalTimeSeriesIntraday
// Returns the active zone minutes intraday data for a 24 hour period by specifying a date range and/or time range.
func GetAzmbyIntervalTimeSeriesIntraday(client *http.Client, startDate time.Time, endDate time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/%s/time/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}

// GetAzmbyDateTimeSeriesIntraday
// Returns the active zone minutes intraday data for a 24 hour period by specifying a date and/or time range.
func GetAzmbyDateTimeSeriesIntraday(client *http.Client, date time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/1d/%s/time/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}

// GetAzmbyDateIntraday
// Returns the active zone minutes intraday data for a 24 hour period by specifying a date and/or time range.
func GetAzmbyDateIntraday(client *http.Client, date time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/1d/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}

// GetAzmbyIntervalIntraday
// Returns the active zone minutes intraday data for a 24 hour period by specifying a date range and/or time range.
func GetAzmbyIntervalIntraday(client *http.Client, startDate time.Time, endDate time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/active-zone-minutes/date/%s/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}
