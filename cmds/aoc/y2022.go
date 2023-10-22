package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/jbert/aoc"
	year "github.com/jbert/aoc/y2022"
)

func main() {
	args := os.Args

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		args = args[1:]
	}

	if len(args) != 3 {
		log.Fatalf("Must have at exactly 2 args")
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Couldn't parse [%d] as number: %s", args[1], err)
	}

	test := true
	if args[2] == "false" {
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
	case 2:
		d = year.NewDay2()
	case 3:
		d = year.NewDay3()
	case 4:
		d = year.NewDay4()
	case 5:
		d = year.NewDay5()
	case 8:
		d = year.NewDay8()
	case 9:
		d = year.NewDay9()
	case 10:
		d = year.NewDay10()
	case 11:
		d = year.NewDay11()
	case 12:
		d = year.NewDay12()
	case 13:
		d = year.NewDay13()
	case 14:
		d = year.NewDay14()
	case 15:
		d = year.NewDay15()
	case 16:
		d = year.NewDay16()
	case 17:
		d = year.NewDay17()
	default:
		return nil, fmt.Errorf("Unknown day [%d]", day)
	}
	return d, nil
}
