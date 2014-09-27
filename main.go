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

	area := 0
	areaStarted := false

	for level >= 0 {
		areaStarted = false
		areaSection := 0

		for i := range vals {
			test := vals[i]
			if test >= level {
				areaStarted = true
				area += areaSection
				areaSection = 0
			} else if (areaStarted) {
				areaSection += 1
			}
		}	
		level -= 1
	}

	return area
}

func main() {
	vals := getIntVals(getArgs())
	area := getArea(vals)
	fmt.Println(area)
}