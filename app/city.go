package main

import (
	"fmt"
)

type City struct {
	Name string
	North *City
	West *City
	East *City
	South *City
	Alien *Alien
}

func (city *City) getDirections() []*City {
	directions := make([]*City, 0)
	if city.North != nil {
		directions = append(directions, city.North)
	}
	if city.West != nil {
		directions = append(directions, city.West)
	}
	if city.East != nil {
		directions = append(directions, city.East)
	}
	if city.South != nil {
		directions = append(directions, city.South)
	}
	return directions
}

func (city *City) Destroy(invader *Alien) *Alien {
	if city.North != nil {
		city.North.South = nil
	}
	if city.West != nil {
		city.West.East = nil
	}
	if city.East != nil {
		city.East.West = nil
	}
	if city.South != nil {
		city.South.North = nil
	}
	resident := city.Alien
	fmt.Printf("City %s was destroyed by alien %s and alien %s\n", city.Name, invader.Name, resident.Name)
	city.Alien = nil
	return resident
}
