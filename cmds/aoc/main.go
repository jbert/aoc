package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jbert/aoc"
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

	err = aoc.Run(day, test, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
}
