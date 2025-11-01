/*
 * File: time_components.go
 * Author: Franz Zbinden
 * Date: 09/08/2025
 * Purpose: This program computes and displays the components (hours, minutes,
 *          and seconds) of a moment in time.
 */

package main

import "fmt"

// Computes the components (hours, minutes, and seconds) of a moment in time.
func timeComponents(elapsedSecs int) (hrs, mins, secs int) {

	const (
		secsPerHour, minsPerHour = 3600, 60
	)
	hrs = elapsedSecs / secsPerHour
	remSecs := elapsedSecs % secsPerHour
	mins = remSecs / minsPerHour
	secs = remSecs - mins*minsPerHour

	return hrs, mins, secs

}

// Starts the execution of the program.
func main() {
	elapsed := 10000
	hrs, mins, secs := timeComponents(elapsed)
	fmt.Printf("\nThe time is %d hours, %d minutes and %d seconds", hrs, mins, secs)
}
