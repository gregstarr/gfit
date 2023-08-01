// Electrocardiogram
// The Electrocardiogram (also known as ECG) endpoint is used for querying the user's on-device ECG readings.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/electrocardiogram/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetEcgLogList
// This endpoint is used for querying the user's on-device ECG readings.
func GetEcgLogList(client *http.Client, beforeDate time.Time, afterDate time.Time, sort string, offset int, limit int) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/ecg/list.json&beforeDate=%s&afterDate=%s&sort=%s&offset=%d&limit=%d",
		Host,
		beforeDate.Format("yyyy-MM-dd"),
		afterDate.Format("yyyy-MM-dd"),
		sort,
		offset,
		limit,
	)
	return makeRequest(client, url)
}
