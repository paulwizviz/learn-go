package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type data struct {
	ID    int
	Field string
	Date  time.Time
}

func monthToInt(month string) time.Month {
	if month == "Jan" {
		return time.January
	}
	if month == "Feb" {
		return time.February
	}
	if month == "Mar" {
		return time.March
	}
	if month == "Apr" {
		return time.April
	}
	if month == "May" {
		return time.May
	}
	if month == "Jun" {
		return time.June
	}
	if month == "Jul" {
		return time.July
	}
	if month == "Aug" {
		return time.August
	}
	if month == "Sep" {
		return time.September
	}
	if month == "Oct" {
		return time.October
	}
	if month == "Nov" {
		return time.December
	}
	if month == "Dec" {
		return time.December
	}
	return time.Month(0)
}

func dateToTime(date string) (time.Time, error) {

	d := strings.Split(date, "-")
	yr, err := strconv.Atoi(d[2])
	if err != nil {
		return time.Time{}, err
	}
	mth := monthToInt(d[1])
	day, _ := strconv.Atoi(d[0])
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Date(yr, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

func parse(reader *csv.Reader) (c chan data) {
	c = make(chan data)
	go func() {
	loop:
		for {
			rec, err := reader.Read()
			if err == io.EOF {

				break
			}
			id, err := strconv.Atoi(rec[0])
			if err != nil {
				continue loop
			}
			dt, err := dateToTime(rec[2])
			if err != nil {
				continue loop
			}
			d := data{
				ID:    id,
				Field: rec[1],
				Date:  dt,
			}
			c <- d
		}
		close(c)
	}()
	return
}

func filter(d chan data) (even chan data, odd chan data) {
	even = make(chan data)
	odd = make(chan data)
	go func() {
		for v := range d {
			if v.ID%2 == 0 {
				even <- v
			} else {
				odd <- v
			}
		}
	}()
	return
}

func count(even chan data, odd chan data) {
	go func() {
		for v := range even {
			fmt.Println(v)
		}
	}()
	go func() {
		for v := range odd {
			fmt.Println(v)
		}
	}()
}

func main() {

	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	d := parse(reader)
	even, odd := filter(d)
	go func() {
		for v := range even {
			fmt.Println(v)
		}
	}()
	go func() {
		for v := range odd {
			fmt.Println(v)
		}
	}()
	time.Sleep(1 * time.Second)
}
