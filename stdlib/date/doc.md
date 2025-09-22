# Date and Time

The `time` package is part of the standard library.

## Parsing Date Time String

The package include a parsing function, for example:

```go
layout := "02-01-2006"
// dateString  
t, err := time.Parse(layout, dateString)
if err != nil {
  log.Fatalf("Error parsing date: %v", err)
}
```

Here is a working [example](./parsing_test.go)

## Parsing Layout Format

When you use the Parse function to verify you date time string, you **must** use these specific numbers and abbreviations in your layout string. Any other number or text in your layout string that isn't one of these particular "magic values" will be treated as a literal character that must appear exactly as is in the date/time string you are parsing or formatting.

For example, if your input string is 2023-10-26:

`2006-01-02` is the correct layout because `2006` is the year placeholder, `01` is the month placeholder, and `02` is the day placeholder. The hyphens - are literal separators.

If you tried `2023-10-26` as a layout, Go would look for the literal string "2023-10-26" in the input, which is not what you want for parsing varying dates. You're asking for the "magic numbers" that Go uses as placeholders in its time package layouts. These are derived from the reference time:
Mon Jan 2 15:04:05 MST 2006
Here's the full list of these specific values and what they represent in a layout string:

| Reference Value | Represents | Example |
| --- | --- | --- |
| 2006 | Full year |  2023 |
| 06 | Two-digit year | 23 |
| 01 | Numeric month with leading zero | 01 for January, 10 for October |
| 1 | Numeric month without leading zero | 1 for January, 10 for October  |
| Jan | Abbreviated month name | Jan, Oct |
| January | Full month name | January, October |
| 02 | Day of the month with leading zero | 02, 26 |
| 2 | Day of the month without leading zero | 2, 26 |
| _2 | Day of the month, space-padded for single digits | 2, 26 |
| Mon | Abbreviated weekday name |  Mon, Thu |
| Monday | Full weekday name | Monday, Thursday |
| 15 | Hour in 24-hour format | 03, 15 |
| 03 | Hour in 12-hour format with leading zero | 03, 11 |
| 3 | Hour in 12-hour format without leading zero | 3, 11 |
| 04 | Minute with leading zero | 04, 30 |
| 4 | Minute without leading zero | 4, 30 |
| 05 | Second with leading zero | 05, 59 |
| 5 | Second without leading zero | 5, 59 |
| .000 | Milliseconds | fixed 3 digits |
| .000000 | Microseconds | fixed 6 digits |
| .000000000 | Nanoseconds | fixed 9 digits |
| .999 | Milliseconds | trailing zeros removed |
| .999999 | Microseconds | trailing zeros removed |
| .999999999 | Nanoseconds | trailing zeros removed |
| PM | AM/PM | uppercase |
| pm | am/pm | lowercase |
| MST | Time zone abbreviation | EST, PST, GMT, BST |
| -0700 | Numeric time zone offset | -0700 for 7 hours behind UTC |
| -07:00 | Numeric time zone offset with colon | -07:00 |
| -07 | Numeric time zone offset | hours only |
| Z0700 | ISO 8601 time zone | Z for UTC, otherwise ±hhmm |
| Z07:00 | ISO 8601 time zone | Z for UTC, otherwise ±hh:mm |
| Z07 | ISO 8601 time zone | Z for UTC, otherwise ±hh |

### The trick to remembering

The numbers 1, 2, 3, 4, 5, 6, 7 appear sequentially in the reference time when written in a particular order (e.g., 01/02 03:04:05PM '06 -0700).

1 for Month (January)
2 for Day (the 2nd)
3 for Hour (3 PM in 12-hour format)
4 for Minute (04)
5 for Second (05)
6 for Year (2006)
7 for Time zone offset (MST is -0700)
Any other characters in your layout string that are not one of these specific reference values will be treated as literal characters that must appear exactly as they are in the input/output string.

Here are [working examples](./layout_test.go)
