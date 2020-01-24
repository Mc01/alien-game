package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const filename string = "./maps/map.txt"

func getNumberOfAliens(arguments []string) int {
	/*
		Expects array with single element as input
		Expects that element is an integer
		Converts element to integer and returns as number of aliens
	*/
	if len(arguments) != 1 {
		panic("Please specify < Number of aliens > as input argument!\n")
	} else {
		numberOfAliens, err := strconv.Atoi(arguments[0])
		if err != nil {
			panic(err)
		}
		if numberOfAliens < 1 {
			panic("Number of aliens should be greater than 0")
		}
		return numberOfAliens
	}
}

func createCity(line string, cities map[string]*City) {
	items := strings.Split(line, " ")
	if len(items) < 2 {
		panic("Line should contain at least city name and one connection!\n")
	}

	cityName := items[0]
	city := City{Name: cityName}
	for i := 1; i < len(items); i++ {
		connection := strings.Split(items[i], "=")
		direction := connection[0]
		neighbourName := connection[1]

		neighbour, exists := cities[neighbourName]
		if !exists {
			continue
		}

		switch direction {
		case "north":
			city.North = neighbour
			neighbour.South = &city
		case "west":
			city.West = neighbour
			neighbour.East = &city
		case "east":
			city.East = neighbour
			neighbour.West = &city
		case "south":
			city.South = neighbour
			neighbour.North = &city
		}
	}
	cities[cityName] = &city
}

func randomCity(cities map[string]*City, freeChoices []reflect.Value) (*City, int) {
	chosenPosition := rand.Intn(len(freeChoices))
	cityName := freeChoices[chosenPosition].String()
	city := cities[cityName]
	return city, chosenPosition
}

func createAlien(
	cities map[string]*City,
	freeChoices []reflect.Value,
	aliens map[string]*Alien,
) []reflect.Value {
	// Get random alien name
	alienName := ""
	for {
		alienName = getRandomName()
		if aliens[alienName] == nil {
			break
		}
	}
	fmt.Printf("Random alien name: %s\n", alienName)

	// Pick random city
	city, chosenPosition := randomCity(cities, freeChoices)
	freeChoices = append(freeChoices[:chosenPosition], freeChoices[chosenPosition+1:]...)
	fmt.Printf("Random city: %s\n", city.Name)

	// Create alien
	alien := Alien{Name:alienName, Position:city}
	city.Alien = &alien
	aliens[alien.Name] = &alien

	return freeChoices
}

func main() {
	// Structures for cities and alien
	cities := map[string]*City{}
	aliens := map[string]*Alien{}

	// Get number of Aliens from CLI
	numberOfAliens := getNumberOfAliens(os.Args[1:])
	fmt.Printf("Number of aliens: %d\n", numberOfAliens)

	// Open file
	file := openFile(filename)
	defer closeFile(file)

	// Read and create cities from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		createCity(line, cities)
	}

	cityNames := reflect.ValueOf(cities).MapKeys()
	numberOfCities := len(cityNames)
	if numberOfAliens > numberOfCities {
		panic("Number of aliens has to be equal of less than number of cities\n")
	}

	// Create alien on random cities
	freeChoices := cityNames
	for i := 0; i < numberOfAliens; i++ {
		freeChoices = createAlien(cities, freeChoices, aliens)
		fmt.Printf("City that are left: %s\n", freeChoices)
	}

	fmt.Printf("\n")
	fmt.Printf("Starting the game with: %d aliens\n", numberOfAliens)
	fmt.Printf("\n")

	// Make alien travel and fight
	limit := 10000
	for i := 0; i < limit; i++ {
		fmt.Printf("Round: %d\n", i + 1)
		if len(reflect.ValueOf(cities).MapKeys()) == 0 {
			panic("There are no cities to destroy left\n")
		}
		aliensAlive := reflect.ValueOf(aliens).MapKeys()
		if len(aliensAlive) == 0 {
			panic("There are no aliens alive anymore\n")
		}
		cannotMove := 0
		for j := 0; j < len(aliensAlive); j++ {
			alienName := aliensAlive[j].String()
			alien := aliens[alienName]
			if alien == nil {
				continue
			}
			outcome, chosenCity := alien.Move()
			switch outcome {
			case CannotMove:
				cannotMove += 1
				fmt.Printf("Alien %s cannot move\n", alien.Name)
			case CanMove:
				fmt.Printf("Alien %s was moved to city %s\n", alien.Name, chosenCity.Name)
			case Fight:
				resident := chosenCity.Destroy(alien)
				delete(cities, chosenCity.Name)
				delete(aliens, alien.Name)
				delete(aliens, resident.Name)
			}
		}
		aliensSurvived := len(reflect.ValueOf(aliens).MapKeys())
		if aliensSurvived == 0 {
			fmt.Printf("All aliens were destroyed\n")
			break
		} else if aliensSurvived == cannotMove {
			fmt.Printf("Aliens cannot move anymore\n")
			break
		}

		if i == limit - 1 {
			fmt.Printf("Ending iteration with %d rounds\n", limit)
		}
	}
}
