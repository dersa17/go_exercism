package cars

// CalculateWorkingCarsPerHour calculates how many working cars are produced per hour,
// given the production rate (carsPerHour) and the success rate in percent.
func CalculateWorkingCarsPerHour(carsPerHour int, successRate float64) float64 {
	return float64(carsPerHour) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are produced per minute,
// given the production rate (carsPerHour) and the success rate in percent.
func CalculateWorkingCarsPerMinute(carsPerHour int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(carsPerHour, successRate) / 60)
}


// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10      // Number of full groups of 10 cars
	individualCars := carsCount % 10   // Remaining cars that don't form a full group
	return uint(groupsOfTen*95000 + individualCars*10000)
}

