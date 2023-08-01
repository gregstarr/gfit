// Heart Rate Time Series
// The Heart Rate Time Series endpoints are used for querying the user's heart rate data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/heartrate-timeseries/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetHeartByDateRange
// Returns the time series data in the specified range for a given resource in the format requested using units in the unit systems that corresponds to the Accept-Language header provided.
func GetHeartByDateRange(client *http.Client, baseDate string, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/%s.json",
		Host,
		baseDate,
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetHeartByDatePeriod
// Returns the time series data in the specified range for a given resource in the format requested using units in the unit systems that corresponds to the Accept-Language header provided.
func GetHeartByDatePeriod(client *http.Client, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/heart/date/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}
