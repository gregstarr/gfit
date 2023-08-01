// Subscriptions
// The Subscription endpoints allow an application to request notifications to user's specific data collection.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/subscription/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
)

// GetSubscriptionsList
// Retreives a list of a user's subscriptions for your application in the format requested. You can either fetch subscriptions for a specific collection or the entire list of subscriptions for the user. For best practice, make sure that your application maintains this list on your side and use this endpoint only to periodically ensure data consistency.
func GetSubscriptionsList(client *http.Client, collectionPath string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/%s/apiSubscriptions.json",
		Host,
		collectionPath,
	)
	return makeRequest(client, url)
}
