package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/y2022"
	"github.com/jbert/aoc/y2025"
	"github.com/jbert/aoc/year"
)

func main() {

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write mem profile to file")
	var numYear = flag.Int("year", time.Now().Year(), "year to run (default current year)")
	flag.Parse()
	args := flag.Args()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		fmt.Printf("Wrote heap profile")
	}

	if len(args) != 2 {
		log.Fatalf("Must have at exactly 2 args, got [%v]", args)
	}
	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Couldn't parse [%s] as number: %s", args[1], err)
	}

	test := true
	if args[1] == "false" {
		test = false
	}

	years := map[int]year.Year{
		2022: y2022.Y,
		2025: y2025.Y,
	}
	year, ok := years[*numYear]
	if !ok {
		log.Fatalf("can't find year [%d]", *numYear)
	}

	err = aoc.Run(year, day, test, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
}
