package birdwatcher

func sumSlice(s []int) int {
	var t int
	for _, v := range s {
		t += v
	}
	return t
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	return sumSlice(birdsPerDay)
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	return sumSlice(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
