// Cardio Fitness Score (VO2 Max)
// The Cardio Fitness Score (also known as VO2 Max) endpoints are used for querying the maximum or optimum rate at which the user’s heart, lungs, and muscles can effectively use oxygen during exercise.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/cardio-fitness-score/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetVo2MaxSummaryByDate
// This endpoint returns the Cardio Fitness Score (VO2 Max) data for a single date. VO2 Max values will be shown as a range if no run data is available or a single numeric value if the user uses a GPS for runs.
func GetVo2MaxSummaryByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/cardioscore/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetVo2MaxSummaryByInterval
// This endpoint returns the Cardio Fitness Score (VO2 Max) data for a date range. VO2 Max values will be shown as a range if no run data is available or a single numeric value if the user uses a GPS for runs.
func GetVo2MaxSummaryByInterval(client *http.Client, startDate time.Time, endDate time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/cardioscore/date/%s/%s.json",
		Host,
		startDate.Format("yyyy-MM-dd"),
		endDate.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}
