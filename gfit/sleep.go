// Sleep
// The Sleep endpoints are used for querying and modifying the user’s general sleep data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/sleep/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetSleepByDate
// The Get Sleep Logs by Date endpoint returns a summary and list of a user's sleep log entries (including naps) as well as detailed sleep entry data for a given day.
func GetSleepByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.2/user/-/sleep/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetSleepByDateRange
// The Get Sleep Logs by Date Range endpoint returns a list of a user's sleep log entries (including naps) as well as detailed sleep entry data for a given date range (inclusive of start and end dates).
func GetSleepByDateRange(client *http.Client, baseDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.2/user/-/sleep/date/%s/%s.json",
		Host,
		baseDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetSleepList
// The Get Sleep Logs List endpoint returns a list of a user's sleep logs (including naps) before or after a given day with offset, limit, and sort order.
func GetSleepList(client *http.Client, beforeDate time.Time, afterDate time.Time, sort string, offset int, limit int) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.2/user/-/sleep/list.json&beforeDate=%s&afterDate=%s&sort=%s&offset=%d&limit=%d",
		Host,
		beforeDate.Format("yyyy-MM-dd"),
		afterDate.Format("yyyy-MM-dd"),
		sort,
		offset,
		limit,
	)
	return makeRequest(client, url)
}

// GetSleepGoal
// Returns the user's sleep goal.
func GetSleepGoal(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.2/user/-/sleep/goal.json",
		Host,
	)
	return makeRequest(client, url)
}
