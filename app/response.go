package main

const url = "https://randomuser.me/api/?inc=name&nat=us"
/*
	Example of response
	{
		"results":[{
			"name":{
				"title":"Miss",
				"first":"Eline",
				"last":"Gauthier"
			}
		}],
		"info":{
			"seed":"a8f1889a62987385",
			"results":1,
			"page":1,"
			version":"1.3"
		}
	}
*/

type ResponseUserName struct {
	Title string
	First string
	Last string
}

type ResponseResult struct {
	Name ResponseUserName
}

type ResponsePagination struct {
	Results int
	Page int
}

type Response struct {
	Results []ResponseResult
	Info ResponsePagination
}
