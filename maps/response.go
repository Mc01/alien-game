package main

const url = "https://randomuser.me/api/?inc=location&nat=us"
/*
	Example of response
	{
		"results":[{
			"location":{
				"street":{
					"number":1392,
					"name":"Oak Lawn Ave"
				},
				"city":"Knoxville",
				"state":"Michigan",
				"country":"United States",
				"postcode":12590,
				"coordinates":{
					"latitude":"-84.8664",
					"longitude":"53.1111"
				},
				"timezone":{
					"offset":"-9:00",
					"description":"Alaska"
				}
			}
		}],
		"info":{
			"seed":"423cecade9550d19",
			"results":1,
			"page":1,
			"version":"1.3"
		}
	}
*/

type ResponseLocation struct {
	City string
	Country string
}

type ResponseResult struct {
	Location ResponseLocation
}

type ResponsePagination struct {
	Results int
	Page int
}

type Response struct {
	Results []ResponseResult
	Info ResponsePagination
}
