package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/docopt/docopt-go"
)

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

func main() {
	vals := getIntVals(getArgs())
	fmt.Println(vals)
}