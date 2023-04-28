package codewarsSolutions

import "strconv"

/*
	Write a function, which takes a non-negative integer (seconds) as input and
	returns the time in a human-readable format (HH:MM:SS)
	- HH = hours, padded to 2 digits, range: 00-99
	- MM = minutes, padded to 2 digits, range: 00-59
	- SS = seconds, padded to 2 digits, range: 00-59

	The maximum time never exceeds 359999(99:59:59)
*/
func HumanReadableTime(seconds int) string {
	// get the hours
	hours := seconds / 3600
	remainder := seconds % 3600

	// get the minutes and seconds
	minutes := remainder / 60
	secs := remainder % 60

	time := []int{hours, minutes, secs}
	readableString := ""
	// convert to human format
	for i := 0; i < len(time); i++ {
		currentVal := ""

		if time[i] < 10 {
			currentVal = "0" + strconv.Itoa(time[i])
		} else {
			currentVal = strconv.Itoa(time[i])
		}

		if i < len(time)-1 {
			readableString += currentVal + ":"
		} else {
			readableString += currentVal
		}
	}

	//readableString := hours + ":" + minutes + ":" + secs
	return readableString
}
