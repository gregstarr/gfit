// Heart Rate Variability
// The Heart Rate Variability (HRV) endpoints are used for querying the user’s heart rate variability data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/heartrate-variability/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetHrvSummaryDate
// This endpoint returns the Heart Rate Variability (HRV) data for a single date. HRV data applies specifically to a user’s “main sleep,” which is the longest single period of time asleep on a given date.
func GetHrvSummaryDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/hrv/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetHrvSummaryInterval
// This endpoint returns the Heart Rate Variability (HRV) data for a date range. HRV data applies specifically to a user’s “main sleep,” which is the longest single period of time asleep on a given date.
func GetHrvSummaryInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/hrv/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
