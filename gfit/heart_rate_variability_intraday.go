// Heart Rate Variability Intraday
// The Heart Rate Variability Intraday endpoint returns the HRV rate at various times and returns Root Mean Square of Successive Differences (rmssd), Low Frequency (LF), High Frequency (HF), and Coverage data for a given measurement.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetHrvIntradayByInterval
// This endpoint returns the Heart Rate Variability (HRV) intraday data for a single date. HRV data applies specifically to a user’s “main sleep,” which is the longest single period of time asleep on a given date.
func GetHrvIntradayByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/hrv/date/%s/%s/all.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetHrvIntradayByDate
// This endpoint returns the Heart Rate Variability (HRV) intraday data for a single date. HRV data applies specifically to a user’s “main sleep,” which is the longest single period of time asleep on a given date.
func GetHrvIntradayByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/hrv/date/%s/all.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
