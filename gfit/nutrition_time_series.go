// Nutrition Time Series
// The Nutrition Time Series endpoints are used for querying food- and water-related data over a period of time either by specifying a date and time period, or a date range.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/nutrition-timeseries/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetFoodsByDateRange
// Updates a user's daily activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
func GetFoodsByDateRange(client *http.Client, resourcePath string, baseDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/%s/date/%s/%s.json",
		Host,
		resourcePath,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetFoodsResourceByDatePeriod
// Updates a user's daily activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
func GetFoodsResourceByDatePeriod(client *http.Client, resourcePath string, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/%s/date/%s/%s.json",
		Host,
		resourcePath,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}
