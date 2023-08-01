// Temperature
// The Temperature endpoints are used for querying either the core temperature data logged manually by the user, or the skin temperature recorded by the device while the user is asleep.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/temperature/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetTempCoreSummaryByInterval
// Returns Temperature (Core) data for a date range. Temperature (Core) data applies specifically to data logged manually by the user throughout the day and the maximum date range cannot exceed 30 days.
func GetTempCoreSummaryByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/temp/core/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetTempSkinSummaryDate
// Returns the Temperature (Skin) data for a single date. Temperature (Skin) data applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
func GetTempSkinSummaryDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/temp/skin/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetTempCoreSummaryByDate
// Returns the Temperature (Core) data for a single date. Temperature (Core) data applies specifically to data logged manually by the user throughout the day.
func GetTempCoreSummaryByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/temp/core/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetTempSkinSummaryByInterval
// Returns Temperature (Skin) data for a date range. It only returns a value for dates on which the Fitbit device was able to record Temperature (skin) data and the maximum date range cannot exceed 30 days.
func GetTempSkinSummaryByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/temp/skin/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
