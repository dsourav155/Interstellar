package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("You have", fuel, "gallons of fuel left.")
}
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Merury":
		fuel = 500000
	case "Venus":
		fuel = 300000
	case "Mars":
		fuel = 700000
	case "Jupiter":
		fuel = 1000000
	case "Saturn":
		fuel = 1400000
	case "Uranus":
		fuel = 1700000
	case "Neptune":
		fuel = 2200000
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome to the", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	fuel := 1000000
	planetChoice := "Neptune"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
