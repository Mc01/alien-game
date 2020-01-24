package main

func getRandomUser() ResponseResult {
	data := getUrl(url)
	return data.Results[0]
}

func getRandomName() string {
	randomUser := getRandomUser()
	return randomUser.Name.First
}
