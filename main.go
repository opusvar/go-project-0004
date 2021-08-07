package main

import f "fmt"

func fuelGuage(fuel int) int {
	f.Println("Remaining Fuel:", fuel)
    return fuel
}

// adapt to format input string to title case

func calculateFuel(planet string) int {
	var fuelNeeded int
    switch planet {
	    case "Venus":
			fuelNeeded = 300000
		case "Mercury":
			fuelNeeded = 500000
		case "Mars":
			fuelNeeded = 700000
		default:
			fuelNeeded = 0
			f.Print("Planet not found")

	}
	f.Println("Calculating fuel needs for trip to", planet)
	f.Println("Fuel needed for trip", fuelNeeded)
    return fuelNeeded
}

func greetPlanet(planet string) {
	f.Println("Attention Passengers and Crew.")
	f.Println("We've arrived at", planet)

}

func cantFly() {
	f.Println("We donot have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int{
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

func main(){
    var fuel int
	var planetChoice string
    
	fuel = 1000000
	planetChoice = "Venus"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGuage(fuel)



}

// If you want to challenge yourself more:

// Add a variable that keeps track of which planet your spaceship is on.
// Create a function that returns your spaceship back to your home planet.
// Add more destinations and allow for traveling between different planets.