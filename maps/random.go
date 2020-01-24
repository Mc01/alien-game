package main

import "strings"

func getRandomLocation() ResponseResult {
	data := getUrl(url)
	return data.Results[0]
}

func getRandomCityName() string {
	randomLocation := getRandomLocation()
	cityName := randomLocation.Location.City
	cityName = strings.ReplaceAll(cityName, " ", "")
	cityName = strings.ReplaceAll(cityName, "-", "")
	cityName = strings.ReplaceAll(cityName, ".", "")
	return cityName
}
