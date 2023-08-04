package changedateformat

import (
    "time"
    "fmt"
)

func ChangeDateFormat(dateString string) string{
	//parse the string to the format YYYY/MM/DD
	parsedDateString :=  dateString[:4] + "-" + dateString[5:7] + "-" + dateString[8:]
    // Create a time.Time object from the date string
    date, err := time.Parse("2006-01-02", parsedDateString)
    if err != nil {
        // Return an error if parsing fails
        fmt.Println(err)
    }
	// parse to the rfc3339 format
	rfc3339Date := date.Format(time.RFC3339)
    // Return the new date string
    return rfc3339Date
}
