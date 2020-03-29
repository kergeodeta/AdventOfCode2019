package main

import (
	"aoc/day1"
	"aoc/day2"
	"flag"
)

var day int

func init() {
	flag.IntVar(&day, "day", 1, "Which day would you to run")
	flag.Parse()
}

func main() {
	switch day {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	}
}
