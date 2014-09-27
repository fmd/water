package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/docopt/docopt-go"
)

/**
 * Returns command-line arguments as strings.
 */
func getArgs() string {
	usage := `Water.

	Usage:
	  water <values>
	  water -h | --help
	  water --version

	Options:
	  -h --help     Show this screen.
	  --version     Show version.`

	arguments, err := docopt.Parse(usage, nil, true, "Water 0.0", false)
	if err != nil {
		panic(err)
	}

	return arguments["<values>"].(string)
}

/**
 * Returns an int array from comma separated string of integers.
 */
func getIntVals(args string) []int {
	var intvals []int
	strvals := strings.Split(args, ",")
	for i := range strvals {
		val, err := strconv.Atoi(strvals[i])
		if err != nil {
			panic(err)
		}

		intvals = append(intvals, val)
	}

	return intvals
}

/**
 * Returns the max value in the slice.
 */
func max(vals []int) int {
	r := 0
	for i := range vals {
		if vals[i] > r {
			r = vals[i]
		}
	}
	return r
}

/** 
 * Takes slice of ints and returns area based on problem spec.
 */
func getArea(vals []int) int {

	//We go from top to bottom to find the area.
	level := max(vals)

	//Set the area to 0
	area := 0
	areaStarted := false

	//Loop from top to bottom of problem. We don't need to test at 0.
	for level > 0 {

		//At the start of each level, no section has been started.
		areaStarted = false
		areaSection := 0

		//Loop through the bars. 
		for i := range vals {

			//Get the height of the current bar.
			test := vals[i]

			//If the height of the current bar reaches the level we're testing:
			if test >= level {

				//Start a new area, and add the previous section's area to the total.
				areaStarted = true
				areaSection = 0
				area += areaSection

			//Otherwise, if an area has been started, add 1 to the current section's area.
			} else if (areaStarted) {
				areaSection += 1
			}
		}

		//Test the next lowest area.
		level -= 1
	}

	//Return the total area.
	return area
}

func main() {

	//Get the array from the command line arguments.
	vals := getIntVals(getArgs())

	//Run the algorithm.
	area := getArea(vals)

	//Print the area.
	fmt.Println(area)
}