package main

import f "fmt"

func fuelGuage(fuel int) int{
	f.Println("Remaining Fuel:", fuel)
    return fuel
}

func main(){
    fuelGuage(30)
}