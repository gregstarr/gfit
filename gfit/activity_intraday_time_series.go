// Activity Intraday Time Series
// The Activity Intraday Time Series endpoints retrieve the data for a given resource on a specific date or 24 hour period.  Intraday support can extend the detail-level response to include 1min, 5min and 15min for the activity.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/intraday/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetActivitiesResourceByDateRangeTimeSeriesIntraday
// Returns the Intraday Time Series for a given resource in the format requested.
func GetActivitiesResourceByDateRangeTimeSeriesIntraday(client *http.Client, resourcePath string, date time.Time, endDate time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/%s/date/%s/%s/%s/time/%s/%s.json",
		Host,
		resourcePath,
		date.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}

// GetActivitiesResourceByDateRangeIntraday
// Returns the Activity Intraday Time Series for a given resource in the format requested.
func GetActivitiesResourceByDateRangeIntraday(client *http.Client, resourcePath string, baseDate time.Time, endDate time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/%s/date/%s/%s/%s.json",
		Host,
		resourcePath,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}

// GetActivitiesResourceByDateTimeSeriesIntraday
// Returns the Intraday Time Series for a given resource in the format requested.
func GetActivitiesResourceByDateTimeSeriesIntraday(client *http.Client, resourcePath string, date time.Time, detailLevel string, startTime string, endTime string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/%s/date/%s/1d/%s/time/%s/%s.json",
		Host,
		resourcePath,
		date.Format("yyyy-MM-dd"),
		detailLevel,
		startTime,
		endTime,
	)
	return makeRequest(client, url)
}

// GetActivitiesResourceByDateIntraday
// Returns the Intraday Time Series for a given resource in the format requested.
func GetActivitiesResourceByDateIntraday(client *http.Client, resourcePath string, date time.Time, detailLevel string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/%s/date/%s/1d/%s.json",
		Host,
		resourcePath,
		date.Format("yyyy-MM-dd"),
		detailLevel,
	)
	return makeRequest(client, url)
}
