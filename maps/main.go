package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

func getLimitOfCities(arguments []string) int {
	/*
		Expects array with single element as input
		Expects that element is an integer
		Converts element to integer and returns as limit of cities
	*/
	if len(arguments) != 1 {
		panic("Please specify < Limit of cities > as input argument!\n")
	} else {
		numberOfCities, err := strconv.Atoi(arguments[0])
		if err != nil {
			panic(err)
		}
		if numberOfCities < 1 {
			panic("Number of cities should be greater than 0")
		}
		return numberOfCities
	}
}

func createRootCity(cities map[string]*City) {
	cityA := City{Name:getRandomCityName()}
	cities[cityA.Name] = &cityA
}

func addCity(cities map[string]*City) bool {
	cityName := ""
	for {
		cityName = getRandomCityName()
		if cities[cityName] == nil {
			break
		}
	}
	existingCities := reflect.ValueOf(cities).MapKeys()
	randomPosition := rand.Intn(len(existingCities))
	rootCity := cities[existingCities[randomPosition].String()]
	directions := rootCity.getFreeDirections()
	if len(directions) > 0 {
		newCity := City{Name:cityName}
		randomDirection := rand.Intn(len(directions))
		switch directions[randomDirection] {
		case "north":
			rootCity.North = &newCity
			newCity.South = rootCity
		case "west":
			rootCity.West = &newCity
			newCity.East = rootCity
		case "east":
			rootCity.East = &newCity
			newCity.West = rootCity
		case "south":
			rootCity.South = &newCity
			newCity.North = rootCity
		}
		cities[newCity.Name] = &newCity
		return true
	} else {
		return false
	}
}

func main() {
	cities := map[string]*City{}

	limit := getLimitOfCities(os.Args[1:])
	fmt.Printf("Will try to create cities until reach limit: %d\n", limit)

	createRootCity(cities)
	for i := 0; i < limit; i++ {
		added := addCity(cities)
		if added {
			fmt.Printf("%d - Success. Added new city!\n", i)
		} else {
			fmt.Printf("%d - Fail. Randomly picked chosen position.\n", i)
		}
	}

	file := createFile("maps/map.txt")
	defer closeFile(file)

	writer := openWriter(file)

	citiesToParse := reflect.ValueOf(cities).MapKeys()
	for i := 0; i < len(citiesToParse); i++ {
		cityName := citiesToParse[i].String()
		writeLine(writer, cities[cityName].getLineFromCity())
	}

	fmt.Printf("Map generated!")
}
