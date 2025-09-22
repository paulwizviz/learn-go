package date

import (
	"fmt"
	"time"
)

func Example_predefinedLayoout() {

	// A sample time to format for demonstration purposes
	sampleTime := time.Date(2025, time.May, 24, 11, 0, 25, 123456789, time.Local) // Using 2025-05-24 11:00:25.123456789 (your local time zone)

	// Using Go's predefined layouts (highly recommended for standard formats)
	fmt.Printf("RFC3339:   %s (Layout: time.RFC3339)\n", sampleTime.Format(time.RFC3339))
	fmt.Printf("ANSIC:     %s (Layout: time.ANSIC)\n", sampleTime.Format(time.ANSIC))
	fmt.Printf("UnixDate:  %s (Layout: time.UnixDate)\n", sampleTime.Format(time.UnixDate))
	fmt.Printf("Stamp:     %s (Layout: time.Stamp)\n", sampleTime.Format(time.Stamp))

	// Output:
	// RFC3339:   2025-05-24T11:00:25+01:00 (Layout: time.RFC3339)
	// ANSIC:     Sat May 24 11:00:25 2025 (Layout: time.ANSIC)
	// UnixDate:  Sat May 24 11:00:25 BST 2025 (Layout: time.UnixDate)
	// Stamp:     May 24 11:00:25 (Layout: time.Stamp)
}

func Example_customlayout() {

	// A sample time to format for demonstration purposes
	sampleTime := time.Date(2025, time.May, 24, 11, 0, 25, 123456789, time.Local) // Using 2025-05-24 11:00:25.123456789 (your local time zone)

	// 1. Common Date Formats
	fmt.Println("1. Common Date Formats:")
	// YYYY-MM-DD
	layout1 := "2006-01-02"
	fmt.Printf("  YYYY-MM-DD:         %s (Layout: \"%s\")\n", sampleTime.Format(layout1), layout1)

	// MM/DD/YYYY
	layout2 := "01/02/2006"
	fmt.Printf("  MM/DD/YYYY:         %s (Layout: \"%s\")\n", sampleTime.Format(layout2), layout2)

	// DD-MM-YYYY
	layout3 := "02-01-2006"
	fmt.Printf("  DD-MM-YYYY:         %s (Layout: \"%s\")\n", sampleTime.Format(layout3), layout3)

	// Month Day, Year (e.g., May 24, 2025)
	layout4 := "Jan 2, 2006"
	fmt.Printf("  Month Day, Year:    %s (Layout: \"%s\")\n", sampleTime.Format(layout4), layout4)

	// Full Month Name Day, Year (e.g., May 24, 2025)
	layout5 := "January 2, 2006"
	fmt.Printf("  Full Month, Day, Y: %s (Layout: \"%s\")\n", sampleTime.Format(layout5), layout5)

	// Weekday, Full Month Day, Year (e.g., Saturday, May 24, 2025)
	layout6 := "Monday, January 2, 2006"
	fmt.Printf("  Weekday, Full M, D, Y: %s (Layout: \"%s\")\n", sampleTime.Format(layout6), layout6)

	fmt.Println("\n2. Time Formats:")
	// HH:MM:SS (24-hour)
	layout7 := "15:04:05"
	fmt.Printf("  HH:MM:SS (24h):     %s (Layout: \"%s\")\n", sampleTime.Format(layout7), layout7)

	// HH:MM AM/PM (12-hour)
	layout8 := "03:04 PM"
	fmt.Printf("  HH:MM AM/PM (12h):  %s (Layout: \"%s\")\n", sampleTime.Format(layout8), layout8)

	// H:M:S AM/PM (12-hour, no leading zeros)
	layout9 := "3:4:5 PM"
	fmt.Printf("  H:M:S AM/PM (12h):  %s (Layout: \"%s\")\n", sampleTime.Format(layout9), layout9)

	fmt.Println("\n3. Combined Date and Time Formats:")
	// YYYY-MM-DD HH:MM:SS
	layout10 := "2006-01-02 15:04:05"
	fmt.Printf("  Combined (24h):     %s (Layout: \"%s\")\n", sampleTime.Format(layout10), layout10)

	// YYYY-MM-DD HH:MM:SS AM/PM
	layout11 := "2006-01-02 03:04:05 PM"
	fmt.Printf("  Combined (12h):     %s (Layout: \"%s\")\n", sampleTime.Format(layout11), layout11)

	fmt.Println("\n4. Time Zones and Precision:")
	// With time zone abbreviation
	layout12 := "2006-01-02 15:04:05 MST"
	fmt.Printf("  With Time Zone (abbr): %s (Layout: \"%s\")\n", sampleTime.Format(layout12), layout12)

	// With numeric time zone offset (e.g., +0100 for BST)
	layout13 := "2006-01-02 15:04:05 -0700"
	fmt.Printf("  With Time Zone (offset): %s (Layout: \"%s\")\n", sampleTime.Format(layout13), layout13)

	// With numeric time zone offset and colon
	layout14 := "2006-01-02 15:04:05 -07:00"
	fmt.Printf("  With Time Zone (offset:colon): %s (Layout: \"%s\")\n", sampleTime.Format(layout14), layout14)

	// With milliseconds (fixed 3 digits)
	layout15 := "2006-01-02 15:04:05.000"
	fmt.Printf("  With Milliseconds:  %s (Layout: \"%s\")\n", sampleTime.Format(layout15), layout15)

	// With nanoseconds (flexible/trailing zeros removed for .999)
	layout16 := "2006-01-02 15:04:05.999999999"
	fmt.Printf("  With Nanoseconds:   %s (Layout: \"%s\")\n", sampleTime.Format(layout16), layout16)

	// Output:
	// 1. Common Date Formats:
	//   YYYY-MM-DD:         2025-05-24 (Layout: "2006-01-02")
	//   MM/DD/YYYY:         05/24/2025 (Layout: "01/02/2006")
	//   DD-MM-YYYY:         24-05-2025 (Layout: "02-01-2006")
	//   Month Day, Year:    May 24, 2025 (Layout: "Jan 2, 2006")
	//   Full Month, Day, Y: May 24, 2025 (Layout: "January 2, 2006")
	//   Weekday, Full M, D, Y: Saturday, May 24, 2025 (Layout: "Monday, January 2, 2006")
	//
	// 2. Time Formats:
	//   HH:MM:SS (24h):     11:00:25 (Layout: "15:04:05")
	//   HH:MM AM/PM (12h):  11:00 AM (Layout: "03:04 PM")
	//   H:M:S AM/PM (12h):  11:0:25 AM (Layout: "3:4:5 PM")
	//
	// 3. Combined Date and Time Formats:
	//   Combined (24h):     2025-05-24 11:00:25 (Layout: "2006-01-02 15:04:05")
	//   Combined (12h):     2025-05-24 11:00:25 AM (Layout: "2006-01-02 03:04:05 PM")
	//
	// 4. Time Zones and Precision:
	//   With Time Zone (abbr): 2025-05-24 11:00:25 BST (Layout: "2006-01-02 15:04:05 MST")
	//   With Time Zone (offset): 2025-05-24 11:00:25 +0100 (Layout: "2006-01-02 15:04:05 -0700")
	//   With Time Zone (offset:colon): 2025-05-24 11:00:25 +01:00 (Layout: "2006-01-02 15:04:05 -07:00")
	//   With Milliseconds:  2025-05-24 11:00:25.123 (Layout: "2006-01-02 15:04:05.000")
	//   With Nanoseconds:   2025-05-24 11:00:25.123456789 (Layout: "2006-01-02 15:04:05.999999999")
}
