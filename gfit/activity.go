// Activity
// The Activity endpoints are used to query and/or modify a Fitbit user's daily activity data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/activity/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetActivitiesLog
// Updates a user's daily activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
func GetActivitiesLog(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetActivitiesTcx
// Retreives the details of a user's location and heart rate data during a logged exercise activity.
func GetActivitiesTcx(client *http.Client, logId string, includePartialTcx bool) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/%s.tcx&includePartialTCX=%t",
		Host,
		logId,
		includePartialTcx,
	)
	return makeRequest(client, url)
}

// GetActivitiesTypes
// Retreives a tree of all valid Fitbit public activities from the activities catelog as well as private custom activities the user created in the format requested.
func GetActivitiesTypes(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/activities.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetActivitiesByDate
// Retrieves a summary and list of a user's activities and activity log entries for a given day.
func GetActivitiesByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetFavoriteActivities
// Returns a list of a user's favorite activities.
func GetFavoriteActivities(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/favorite.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetActivitiesLogList
// Retreives a list of user's activity log entries before or after a given day with offset and limit using units in the unit system which corresponds to the Accept-Language header provided.
func GetActivitiesLogList(client *http.Client, beforeDate time.Time, afterDate time.Time, sort string, offset int, limit int) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/list.json&beforeDate=%s&afterDate=%s&sort=%s&offset=%d&limit=%d",
		Host,
		beforeDate.Format("yyyy-MM-dd"),
		afterDate.Format("yyyy-MM-dd"),
		sort,
		offset,
		limit,
	)
	return makeRequest(client, url)
}

// GetRecentActivities
// Retreives a list of a user's recent activities types logged with some details of the last activity log of that type using units in the unit system which corresponds to the Accept-Language header provided.
func GetRecentActivities(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/recent.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetActivitiesGoals
// Retreives a user's current daily or weekly activity goals using measurement units as defined in the unit system, which corresponds to the Accept-Language header provided.
func GetActivitiesGoals(client *http.Client, period string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/goals/%s.json",
		Host,
		period,
	)
	return makeRequest(client, url)
}

// GetActivitiesTypeDetail
// Returns the detail of a specific activity in the Fitbit activities database in the format requested. If activity has levels, it also returns a list of activity level details.
func GetActivitiesTypeDetail(client *http.Client, activityId string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/activities/%s.json",
		Host,
		activityId,
	)
	return makeRequest(client, url)
}

// GetFrequentActivities
// Retreives a list of a user's frequent activities in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
func GetFrequentActivities(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/activities/frequent.json",
		Host,
	)
	return makeRequest(client, url)
}
