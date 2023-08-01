// Breathing Rate Intraday
// The Breathing Rate Intraday endpoint returns the  average breathing rate throughout the day and categories your breathing rate by sleep stage.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetBreathingRateIntradayByInterval
// This endpoint returns intraday breathing rate data for a date range. It measures the average breathing rate throughout the day and categories your breathing rate by sleep stage. Sleep stages vary between light sleep, deep sleep, REM sleep, and full sleep.
func GetBreathingRateIntradayByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/br/date/%s/%s/all.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetBreathingRateIntradayByDate
// This endpoint returns intraday breathing rate data for a single date. It measures the average breathing rate throughout the day and categories your breathing rate by sleep stage. Sleep stages vary between light sleep, deep sleep, REM sleep, and full sleep.
func GetBreathingRateIntradayByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/br/date/%s/all.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
