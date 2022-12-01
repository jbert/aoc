package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jbert/aoc"
	year "github.com/jbert/aoc/y2022"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Must have exactly 2 args")
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Couldn't parse [%d] as number: %s", os.Args[1], err)
	}

	test := true
	if os.Args[2] == "false" {
		test = false
	}

	d, err := intToDay(day)
	if err != nil {
		log.Fatalf("Can't get day for [%d]: %s", err)
	}

	err = aoc.Run(d, day, test, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
}

func intToDay(day int) (aoc.Day, error) {
	var d aoc.Day

	switch day {
	case 1:
		d = year.NewDay1()
	default:
		return nil, fmt.Errorf("Unknown day [%d]", day)
	}
	return d, nil
}