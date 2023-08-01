// Nutrition
// The Nutrition endpoints are used for querying and modifying the food and water data.
// Full documentation: https://dev.fitbit.com/build/reference/web-api/nutrition/
// code generated 2023-07-31 16:04:19

package gfit

import (
	"fmt"
	"net/http"
	"time"
)

// GetWaterGoal
// Retreives a summary and list of a user's water goal entries for a given day in the requested using units in the unit system that corresponds to the Accept-Language header provided.
func GetWaterGoal(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/water/goal.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetWaterByDate
// Retreives a summary and list of a user's water log entries for a given day in the requested using units in the unit system that corresponds to the Accept-Language header provided.
func GetWaterByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/water/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetFavoriteFoods
// Returns a list of a user's favorite foods in the format requested. A favorite food in the list provides a quick way to log the food via the Log Food endpoint.
func GetFavoriteFoods(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/favorite.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetMeals
// Returns a list of meals created by user in the user's food log in the format requested. User creates and manages meals on the Food Log tab on the website.
func GetMeals(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/meals.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFoodsGoal
// Returns a user's current daily calorie consumption goal and/or foodPlan value in the format requested.
func GetFoodsGoal(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/goal.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFoodsLocales
// Returns the food locales that the user may choose to search, log, and create food in.
func GetFoodsLocales(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/foods/locales.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFoodsByDate
// Retreives a summary and list of a user's food log entries for a given day in the format requested.
func GetFoodsByDate(client *http.Client, date time.Time) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/date/%s.json",
		Host,
		date.Format("yyyy-MM-dd"),
	)
	return makeRequest(client, url)
}

// GetFoodsUnits
// Returns a list of all valid Fitbit food units in the format requested.
func GetFoodsUnits(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/foods/units.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFoodsInfo
// Returns the details of a specific food in the Fitbit food databases or a private food that an authorized user has entered in the format requested.
func GetFoodsInfo(client *http.Client, foodId string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/foods/%s.json",
		Host,
		foodId,
	)
	return makeRequest(client, url)
}

// GetRecentFoods
// Returns a list of a user's frequent foods in the format requested. A frequent food in the list provides a quick way to log the food via the Log Food endpoint.
func GetRecentFoods(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/recent.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFrequentFoods
// Returns a list of a user's frequent foods in the format requested. A frequent food in the list provides a quick way to log the food via the Log Food endpoint.
func GetFrequentFoods(client *http.Client) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/user/-/foods/log/frequent.json",
		Host,
	)
	return makeRequest(client, url)
}

// GetFoodsList
// Returns a list of public foods from the Fitbit food database and private food the user created in the format requested.
func GetFoodsList(client *http.Client, query string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/1/foods/search.json&query=%s",
		Host,
		query,
	)
	return makeRequest(client, url)
}
