package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for i:=0; i < len(birdsPerDay); i++ {
        total = total + birdsPerDay[i]
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    
	total := 0
    start := 7*(week-1)
    for i:=start; i < start+7; i++ {
        total = total + birdsPerDay[i]
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0; i < len(birdsPerDay); i++ {
        if (i % 2 == 0) {
            birdsPerDay[i] = birdsPerDay[i] + 1
        }
    }
    return birdsPerDay
}
