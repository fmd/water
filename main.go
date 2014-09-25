package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Water.

	Usage:
	  water <file>
	  water -h | --help
	  water --version

	Options:
	  -h --help     Show this screen.
	  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Water 0.0", false)
	fmt.Println(arguments)
}