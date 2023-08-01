// Friends
// The Friends endpoints display information about the specified user's friends and the friend's leaderboard.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/friends/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
)

// GetFriends
// Returns data of a user's friends in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
func GetFriends(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.1/user/-/friends.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFriendsLeaderboard
// Returns data of a user's friends in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
func GetFriendsLeaderboard(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1.1/user/-/leaderboard/friends.json",
		Host,
	)
	return makeRequest(client, url)
}
