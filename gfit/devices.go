// Devices
// The Devices endpoints displays information about the devices paired to a user's account.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/devices/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
)

// GetAlarms
// Returns alarms for a device
func GetAlarms(client *http.Client, trackerId int) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/devices/tracker/%d/alarms.json",
		Host,
		trackerId,
	)
	return makeRequest(client, url)
}

// GetDevices
// Returns a list of the Fitbit devices connected to a user's account.
func GetDevices(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/devices.json",
		Host,
	)
	return makeRequest(client, url)
}
