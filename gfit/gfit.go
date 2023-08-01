// auto generated fitbit client library: version 1
// Fitbit provides a Web API for accessing data from Fitbit activity trackers, Aria scale, and manually entered logs. Anyone can develop an application to access and modify a Fitbit user's data on their behalf, so long as it complies with Fitbit Platform Terms of Service. These Swagger UI docs do not currently support making Fitbit API requests directly. In order to make a request, construct a request for the appropriate endpoint using this documentation, and then add an Authorization header to each request with an access token obtained using the steps outlined here: https://dev.fitbit.com/build/reference/web-api/developer-guide/authorization/.
// code generated 2023-07-31 16:04:19

// Package gfit
package gfit

const (
	Host             = "https://api.fitbit.com"
	AuthorizationUrl = "https://www.fitbit.com/oauth2/authorize"
	TokenUrl         = "/oauth2/token"
)

type AuthScope string

const (
	// Location ->GPS and other location data
	Location AuthScope = "location"
	// OxygenSaturation ->SpO2 (Oxygen Saturation) is a measurement of your blood oxygen level
	OxygenSaturation AuthScope = "oxygen_saturation"
	// Settings ->User account and device settings, such as alarms
	Settings AuthScope = "settings"
	// Nutrition ->Calorie consumption and nutrition related features, such as food/water logging, goals, and plans
	Nutrition AuthScope = "nutrition"
	// Sleep ->Sleep logs and related sleep analysis
	Sleep AuthScope = "sleep"
	// Activity ->Activity data and exercise log related features, such as steps, distance, calories burned, and active minutes
	Activity AuthScope = "activity"
	// CardioFitness ->Maximum or optimum rate at which the user’s heart, lungs, and muscles can effectively use oxygen during exercise
	CardioFitness AuthScope = "cardio_fitness"
	// Heartrate ->Continuous heart rate & HRV data and related analysis
	Heartrate AuthScope = "heartrate"
	// Profile ->Basic user information
	Profile AuthScope = "profile"
	// RespiratoryRate ->Respiratory Rate is a measurement of your breaths at night.
	RespiratoryRate AuthScope = "respiratory_rate"
	// Social ->Friend-related features, such as friend list and leaderboard
	Social AuthScope = "social"
	// Temperature ->Core and skin temperature data
	Temperature AuthScope = "temperature"
	// Weight ->Weight and related information, such as body mass index, body fat percentage, and goals
	Weight AuthScope = "weight"
	// Electrocardiogram ->A users on-device ECG readings
	Electrocardiogram AuthScope = "electrocardiogram"
)
