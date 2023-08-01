// Body Time Series
// The Body Time Series endpoints are used for querying body fat and weight data over a period of time either by specifying a date and time period, or a date range.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/body-timeseries/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetBodyResourceByDatePeriod
// Returns time series data in the specified range for a given resource in the format requested using units in the unit system that corresponds to the Accept-Language header provided.
func GetBodyResourceByDatePeriod(client *http.Client, resourcePath string, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/%s/date/%s/%s.json",
		Host,
		resourcePath,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}

// GetBodyResourceByDateRange
// Returns time series data in the specified range for a given resource in the format requested using units in the unit system that corresponds to the Accept-Language header provided.
func GetBodyResourceByDateRange(client *http.Client, resourcePath string, baseDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/%s/date/%s/%s.json",
		Host,
		resourcePath,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
