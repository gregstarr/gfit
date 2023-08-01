// SpO2
// SpO2 (Oxygen Saturation) is a measurement of your blood oxygen level.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/spo2/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetSpO2SummaryByInterval
// This endpoint returns the SpO2 summary data for a date range. SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
func GetSpO2SummaryByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/spo2/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetSpO2SummaryByDate
// This endpoint returns the SpO2 summary data for a single date. SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
func GetSpO2SummaryByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/spo2/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
