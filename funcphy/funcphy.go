package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, initvel, initdist float64) func(float64) float64 {

	funret := func(time float64) float64 {

		return (0.5*acc*math.Pow(time, 2) + initvel*time + initdist)
	}

	return funret
}

func main() {

	var acceleration float64

	fmt.Println("Enter the measure of acceleration: ")
	fmt.Scan(&acceleration)

	var initvelocity float64

	fmt.Println("Enter the measure of initial velocity: ")
	fmt.Scan(&initvelocity)

	var initdistance float64

	fmt.Println("Enter the measure of initial displacement: ")
	fmt.Scan(&initdistance)

	distance := GenDisplaceFn(acceleration, initvelocity, initdistance)

	fmt.Println("The distance paced after 3 seconds: ")
	fmt.Println(distance(3))

	fmt.Println("The distance paced after 5 seconds: ")
	fmt.Println(distance(5))

}
