// Body
// The Body endpoints are used for querying and modifying the user's body fat and weight data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/body/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetBodyFatByDateRange
// Retreives a list of all user's body fat log entries for a given day in the format requested.
func GetBodyFatByDateRange(client *http.Client, baseDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/fat/date/%s/%s.json",
		Host,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetBodyGoals
// Retreives a user's current body fat percentage or weight goal using units in the unit systems that corresponds to the Accept-Language header providedin the format requested.
func GetBodyGoals(client *http.Client, goalType string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/%s/goal.json",
		Host,
		goalType,
	)
	return makeRequest(client, url)
}

// GetWeightByDate
// Retreives a list of all user's body weight log entries for a given day using units in the unit systems which corresponds to the Accept-Language header provided.
func GetWeightByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/weight/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetWeightByDateRange
// Retreives a list of all user's body fat log entries for a given day in the format requested.
func GetWeightByDateRange(client *http.Client, baseDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/weight/date/%s/%s.json",
		Host,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetBodyFatByDatePeriod
// Retreives a list of all user's body fat log entries for a given day in the format requested.
func GetBodyFatByDatePeriod(client *http.Client, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/fat/date/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}

// GetBodyFatByDate
// Retreives a list of all user's body fat log entries for a given day in the format requested.
func GetBodyFatByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/fat/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetWeightByDatePeriod
// Retreives a list of all user's body weight log entries for a given day in the format requested.
func GetWeightByDatePeriod(client *http.Client, date time.Time, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/body/log/weight/date/%s/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
		period,
	)
	return makeRequest(client, url)
}
