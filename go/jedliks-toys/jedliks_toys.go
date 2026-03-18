package jedlik

import "fmt"

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() Car {
	if car.battery < car.batteryDrain {
		return *car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return *car
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(trackDistance int) bool {
	drives := car.battery / car.batteryDrain
	return drives*car.speed >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
