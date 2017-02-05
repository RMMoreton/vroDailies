// Package vroDailies creates daily schedules for Vromans employees.
package vroDailies

// Import...?
import ()

// Keep as many people on the floor as possible, while also scheduling
// breaks as evenly as possible. Everything needs to be in 15 minute
// intervals.

// Take (name, (start, end)) pairs
// Greedy: Calculate when exactly everyone's breaks/lunches *should* be,
// and use that.
//	- Lunch is right in the middle
//		- Only get a lunch if you work more than ?? hours
//	- First break starts at the midpoint between start and lunch
//		- Only get a first break if you work more than ?? hours
//	- Second break starts at the midpoint between lunch and end
//		- Only get a second break if you work more than ?? hours
//	- If first and/or second break don't fall evenly on a quarter hour,
//	  always go earlier.

// These determine obvious properties. Change them as needed.
const FB = 4
const SB = 6
const L = 8

func singlePerson(start, end int) {
	var f, s, l, t int
	t = end - start
	// Do other stuff!
}