package main

import (
	"fmt"
	"math/rand"
)

type Alien struct {
	Name string
	Position *City
}

const(
	CannotMove = 0
	CanMove = 1
	Fight = 2
)

func (alien *Alien) Move() (int, *City) {
	// get movement possibilities
	directions := alien.Position.getDirections()
	possibilities := len(directions)

	// check if alien can move
	if possibilities == 0 {
		return CannotMove, nil
	}

	// pick random position
	choice := rand.Intn(possibilities)

	// move alien
	oldCity := alien.Position
	alien.Position.Alien = nil
	chosenCity := directions[choice]

	fmt.Printf("Moving alien %s from city %s to city %s\n", alien.Name, oldCity.Name, chosenCity.Name)

	// fight resident alien
	if chosenCity.Alien != nil {
		return Fight, chosenCity
	}

	// finish movement
	alien.Position = chosenCity
	alien.Position.Alien = alien

	return CanMove, chosenCity
}
