package date

import (
	"fmt"
	"time"
)

func Example_parsing() {

	// Example: Parsing a date from a string
	dateToParse := "1988-08-16 10:30:00"
	parseLayout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(parseLayout, dateToParse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed \"%s\": %s\n", dateToParse, parsedTime)

	// Example: Parsing a date with full month name
	dateToParse2 := "Friday, December 25, 2020"
	parseLayout2 := "Monday, January 2, 2006"
	parsedTime2, err := time.Parse(parseLayout2, dateToParse2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed \"%s\": %s\n", dateToParse2, parsedTime2)

	// Example: Parsing with different numeric day format (_2 for space padding)
	dateToParse3 := "Apr  5, 2021" // Note the two spaces before 5
	parseLayout3 := "Jan _2, 2006" // Use _2 for space-padded day
	parsedTime3, err := time.Parse(parseLayout3, dateToParse3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed \"%s\": %s\n", dateToParse3, parsedTime3)

	dateToParse4 := "Apr 5, 2021"  // Note the two spaces before 5
	parseLayout4 := "Jan 02, 2006" // Use _2 for space-padded day
	parsedTime4, err := time.Parse(parseLayout4, dateToParse4)
	if err != nil {
		fmt.Printf("To parse: %v Layout: %v\n  Parsed: %v\n  Error: %v", dateToParse4, parseLayout4, parsedTime4, err)
	}

	// Output:
	// Parsed "1988-08-16 10:30:00": 1988-08-16 10:30:00 +0000 UTC
	// Parsed "Friday, December 25, 2020": 2020-12-25 00:00:00 +0000 UTC
	// Parsed "Apr  5, 2021": 2021-04-05 00:00:00 +0000 UTC
	// To parse: Apr 5, 2021 Layout: Jan 02, 2006
	//   Parsed: 0001-01-01 00:00:00 +0000 UTC
	//   Error: parsing time "Apr 5, 2021" as "Jan 02, 2006": cannot parse "5, 2021" as "02"
}
