package parsestring

import (
	"strings"
)

// ParsedInfo represents the parsed information extracted from the input string.
type ParsedInfo struct {
	Date    string
	Hour    string
	Name    string
	Message string
}

// Parsestring takes an input string and parses it to extract relevant information.
func Parsestring(str string) ParsedInfo {
	// Split the input string by the '|' character to separate date/hour and name/message parts.
	substrings := strings.Split(str, "|")

	// Extract the unparsed date and hour from the first part of the split result.
	substringDateAndHourUnparsed := substrings[0]

	// Extract the unparsed name and message from the second part of the split result.
	substringNameAndMessageUnparsed := substrings[1]

	// Split the unparsed date and hour by the space character to get separate date and hour parts.
	substringDateAndHourParsed := strings.Split(substringDateAndHourUnparsed, " ")

	// Split the unparsed name and message by the ':' character to get separate name and message parts.
	substringNameAndMessageParsed := strings.Split(substringNameAndMessageUnparsed, ":")

	// Extract the date and hour from the parsed date/hour part.
	substringDate := substringDateAndHourParsed[0]
	substringHour := substringDateAndHourParsed[1]

	// Extract the name and message from the parsed name/message part.
	substringName := substringNameAndMessageParsed[0]
	substringMessage := substringNameAndMessageParsed[1]

	// Create a new ParsedInfo struct and fill it with the extracted information.
	parsedInfo := ParsedInfo{
		Date:    substringDate,
		Hour:    substringHour,
		Name:    substringName,
		Message: substringMessage,
	}

	// Return the parsed information.
	return parsedInfo
}

