// Heart Rate Intraday Time Series
// The Heart Rate Intraday Time Series retrieves the heart rate data on a specific date or 24 hour period. Intraday support can extend the detail-level response to include 1sec, 1min, 5min or 15min for Heart Rate.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetHeartByDateTimestampIntraday
// Returns the intraday time series for a given resource in the format requested. If your application has the appropriate access, your calls to a time series endpoint for a specific day (by using start and end dates on the same day or a period of 1d), the response will include extended intraday values with a one-minute detail level for that day. Unlike other time series calls that allow fetching data of other users, intraday data is available only for and to the authorized user.
func GetHeartByDateTimestampIntraday(client *http.Client, date time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/1d/%s/time/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}

// GetHeartByDateRangeIntraday
// Returns the intraday time series for a given resource in the format requested. If your application has the appropriate access, your calls to a time series endpoint for a specific day (by using start and end dates on the same day or a period of 1d), the response will include extended intraday values with a one-minute detail level for that day. Unlike other time series calls that allow fetching data of other users, intraday data is available only for and to the authorized user.
func GetHeartByDateRangeIntraday(client *http.Client, date time.Time, endDate time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}

// GetHeartByDateIntraday
// Returns the intraday time series for a given resource in the format requested. If your application has the appropriate access, your calls to a time series endpoint for a specific day (by using start and end dates on the same day or a period of 1d), the response will include extended intraday values with a one-minute detail level for that day. Unlike other time series calls that allow fetching data of other users, intraday data is available only for and to the authorized user.
func GetHeartByDateIntraday(client *http.Client, date time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/1d/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}

// GetHeartByDateRangeTimestampIntraday
// Returns the intraday time series for a given resource in the format requested. If your application has the appropriate access, your calls to a time series endpoint for a specific day (by using start and end dates on the same day or a period of 1d), the response will include extended intraday values with a one-minute detail level for that day. Unlike other time series calls that allow fetching data of other users, intraday data is available only for and to the authorized user.
func GetHeartByDateRangeTimestampIntraday(client *http.Client, date time.Time, endDate time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/%s/%s/time/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}
