package elon

import "fmt"

// Define the 'Drive()' method on Car.
func (c *Car) Drive() *Car {
    if c.batteryDrain <= c.battery {
        c.battery -= c.batteryDrain
        c.distance += c.speed
    }
    return c
}

// Define the 'DisplayDistance() string' method.
func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// Define the 'DisplayBattery() string' method.
func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// Define the 'CanFinish(trackDistance int) bool' method.
func (c *Car) CanFinish(trackDistance int) bool {
    return trackDistance / c.speed * c.batteryDrain <= c.battery
}
