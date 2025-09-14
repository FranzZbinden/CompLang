package main

/*
 * File: time_components.go
 * Author: Franz Zbinden
 * Course: COTI 4039_vh1
 * Porpuse: x
 */

import "fmt" //para formatear input y output

func timeComponents(elapsedSecs int) (hrs, mins, secs int) {
	const (
		secsPerHour   = 3600
		secsPerMinute = 60
	)

	hrs = elapsedSecs / secsPerHour
	remainingSecs := elapsedSecs % secsPerHour

	mins = remainingSecs / secsPerMinute
	secs = remainingSecs % secsPerMinute
	return hrs, mins, secs
}

func main() {
	elapsedSeconds := 12345
	fmt.Println("The number of elapsed seconds is", elapsedSeconds)

	hours, minutes, seconds := timeComponents(elapsedSeconds)
	fmt.Printf("\nThis is %d hours, %d minutes, and %d seconds", hours, minutes, seconds)
}
