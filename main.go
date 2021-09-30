package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
const dateLayout = `2006-01-02`

var fromDateString string
var fromDate time.Time
var toDateString string
var toDate time.Time
var now time.Time

func init() {
	var err error

	now = time.Now().Local()

	flag.StringVar(&fromDateString, "from", "1970-01-01", "the from input date")
	flag.StringVar(&toDateString, "to", "now", "the to input date")
	flag.Parse()

	fromDate, err = time.ParseInLocation(dateLayout, fromDateString, time.Local)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fromDate: %v\n", err)
		os.Exit(1)
	}
	if toDateString == "now" {
		toDate = now
	} else {
		toDate, err = time.ParseInLocation(dateLayout, toDateString, time.Local)
		if err != nil {
			fmt.Fprintf(os.Stderr, "toDate: %v\n", err)
			os.Exit(1)
		}
	}
}

func main() {
	if toDate.Before(fromDate) {
		log.Fatalf("fromDate[%v] before toDate[%v]\n", fromDate, toDate)
	}
	fmt.Printf("from: %v\n", fromDate)
	fmt.Printf("  to: %v\n", toDate)

	dur := toDate.Sub(fromDate)
	fmt.Println(dur)
	fmt.Printf("%d seconds\n", int64(dur.Seconds()))
	fmt.Printf("%d minutes\n", int64(dur.Minutes()))
	hours := int64(dur.Hours())
	fmt.Printf("%d hours\n", hours)
	days := hours / 24
	fmt.Printf("%d days\n", days)
	weeks := days / 7
	fmt.Printf("%d weeks\n", weeks)
	months := monthsBetween(fromDate, toDate)
	fmt.Printf("%d months\n", months)
	years := months / 12
	fmt.Printf("%d years\n", years)
}

func monthsBetween(t0 time.Time, t1 time.Time) int {
	months := 0
	if t0.Equal(t1) {
		return 0
	}
	if t1.Before(t0) {
		t0, t1 = t1, t0
	}
	if t0.After(t1) {
		log.Fatalln("time swap failed")
	}

	var end time.Time = firstOfThisMonth(t1)
	var start time.Time = firstOfThisMonth(t0)
	// fmt.Printf("start: %v\n", start)
	// fmt.Printf("  end: %v\n", end)
	if start.After(end) {
		return 0
	}
	if start.Equal(end) {
		return 0
	}

	years := end.Year() - start.Year()
	months += years * 12
	months += int(end.Month()) - int(start.Month())

	if t0.Day() > t1.Day() {
		months--
	}

	return months
}

func firstOfThisMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}
