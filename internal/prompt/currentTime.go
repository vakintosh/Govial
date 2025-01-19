// This file contains the getCurrentTime function which returns the current time
//  The function returns the hour and minute as integers
// The function is used in the main.go file to get the current time and display it on the right side of the screen
//

package prompt

import (
	"time"
)

func getCurrentTime() (int, int, int) {

	// Get the current time
	currentTime := time.Now()

	// Extract the hour and minute components
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()

	// Return the hour and minute as integers
	return hour, minute, second
}
