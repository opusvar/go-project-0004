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

func main(){
	greetPlanet("Genesis")


}