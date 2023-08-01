// User
// The User endpoints display information about the user's profile information, the regional locale & language settings, and their badges collected
// Full documentation: https://dev.fitbit.com/build/reference/web-api/user/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
)

// GetProfile
// Returns a user's profile. The authenticated owner receives all values. However, the authenticated user's access to other users' data is subject to those users' privacy settings. Numerical values are returned in the unit system specified in the Accept-Language header.
func GetProfile(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/profile.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetBadges
// Retrieves the user's badges in the format requested. Response includes all badges for the user as seen on the Fitbit website badge locker (both activity and weight related.) The endpoint returns weight and distance badges based on the user's unit profile preference as on the website.
func GetBadges(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/badges.json",
		Host,
	)
	return makeRequest(client, url)
}
