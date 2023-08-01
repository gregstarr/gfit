// SpO2 Intraday
// The SpO2 Intraday endpoints return the SpO2 values calculated on a 5-minute exponentially-moving average.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetSpO2IntradayByDate
// This endpoint returns the SpO2 intraday data for a single date. SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
func GetSpO2IntradayByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/spo2/date/%s/all.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetSpO2IntradayByInterval
// This endpoint returns the SpO2 intraday data for a specified date range. SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
func GetSpO2IntradayByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/spo2/date/%s/%s/all.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
