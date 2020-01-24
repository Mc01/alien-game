package main

import "strings"

type City struct {
	Name string
	North *City
	West *City
	East *City
	South *City
}

func concat(a string, b string, sep string) string {
	return strings.Join([]string{a, b}, sep)
}

func createDirection(direction string, city *City) string {
	return concat(direction, city.Name, "=")
}

func joinLineWithDirection(line string, direction string) string {
	return concat(line, direction, " ")
}

func (city *City) getLineFromCity() string {
	line := city.Name
	if city.North != nil {
		direction := createDirection("north", city.North)
		line = joinLineWithDirection(line, direction)
	}
	if city.West != nil {
		direction := createDirection("west", city.West)
		line = joinLineWithDirection(line, direction)
	}
	if city.East != nil {
		direction := createDirection("east", city.East)
		line = joinLineWithDirection(line, direction)
	}
	if city.South != nil {
		direction := createDirection("south", city.South)
		line = joinLineWithDirection(line, direction)
	}
	items := strings.Split(line, " ")
	if len(items) < 2 {
		panic("Cannot create city without directions")
	}
	return line
}

func (city *City) getFreeDirections() []string {
	var directions []string
	if city.North == nil {
		directions = append(directions, "north")
	}
	if city.West == nil {
		directions = append(directions, "west")
	}
	if city.East == nil {
		directions = append(directions, "east")
	}
	if city.South == nil {
		directions = append(directions, "south")
	}
	return directions
}