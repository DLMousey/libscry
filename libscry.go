package main

import (
	"strings"
)

const TDOUBLEQUOTE = "\""
const TSINGLEQUOTE = "'"
const TSPACE = " "
const TOR = "or"
const TAND = "and"
const TTYPE = "t:"
const TTYPEFULL = "type:"
const TCOLOR = "c:"
const TCOLORFULL = "color:"
const TCOLORCORRECTFULL = "colour:"
const TORACLE = "o:"
const TORACLEFULL = "oracle:"

const ETYPE = "type"
const ENAME = "name"
const ECOLOR = "color"
const EORACLE = "oracle"

const CWHITE = "w"
const CBLUE = "u"
const CBLACK = "b"
const CRED = "r"
const CGREEN = "g"
const CCOLORLESS = "c"

var expansions = map[string]string{
	TTYPE:             ETYPE,
	TTYPEFULL:         ETYPE,
	TDOUBLEQUOTE:      ENAME,
	TSINGLEQUOTE:      ENAME,
	TCOLOR:            ECOLOR,
	TCOLORFULL:        ECOLOR,
	TCOLORCORRECTFULL: ECOLOR,
	TORACLE:           EORACLE,
	TORACLEFULL:       EORACLE,
}

var colours = map[string]string{
	CWHITE: "white",
	CBLUE:  "blue",
	CBLACK: "black",
	CRED:   "red",
	CGREEN: "green",
}

type ColumnMapping struct {
	Type    string `yaml:"type"`
	SubType string `yaml:"subtype"`
	Name    string `yaml:"name"`
}

var stopChars []string = []string{
	TSPACE, TOR, TAND, TTYPE, TTYPEFULL, TCOLOR,
	TCOLORFULL, TCOLORCORRECTFULL, TORACLE,
	TORACLEFULL, TDOUBLEQUOTE, TSINGLEQUOTE,
}

func Parse(input string) map[string]string {
	criteria := make(map[string]string)

	// Break string into parts on each space
	// Iterate through each part
	// If a part begins with a quote mark, find the next part that contains another quote to delimit it,
	// then join these parts into a single part and remove the separate ones from the array

	// Some sanitisation first - replace all single quotes with double quotes,
	// strip the newline suffix and lowercase the whole thing
	input = strings.ReplaceAll(input, "'", "\"")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ToLower(input)
	parts := strings.Split(input, TSPACE)

	//lastIdx := 0
	for idx, part := range parts {
		// Iterate over each of the tokens for each part and see if it's contained within this part
		for _, stopChar := range stopChars {
			hastoken := strings.Contains(part, stopChar)

			if hastoken {

				// find the index
				startIndex := strings.Index(part, stopChar)
				endIndex := startIndex + len(stopChar)

				marker := part[startIndex:endIndex]
				expansion := expansions[marker]

				switch expansion {
				case EORACLE:
					extractName(criteria, idx, stopChar, part, parts, expansion)
					break
				case ECOLOR:
					criteria[ECOLOR] = colours[part[endIndex:endIndex+1]]
					break
				case ETYPE:
					criteria[ETYPE] = part[endIndex:]
					break
				default:
					extractName(criteria, idx, stopChar, part, parts, expansion)
					break
				}
			}
		}
	}

	return criteria
}

func extractName(criteria map[string]string, index int, stopChar string, part string, parts []string, expansion string) {
	// If the quote isn't at the beginning of the part, we've probably processed this already
	//if strings.Index(part, TDOUBLEQUOTE) != 0 {
	//	return
	//}

	// First of all check if the current part contains 2 quotes.
	// If there is 2 we'll skip the lookahead and add it straight into the criteria
	quoteCount := strings.Count(part, TDOUBLEQUOTE)
	if quoteCount == 2 {
		criteria[expansion] = part
		return
	}

	// Otherwise we'll have to look through each part in front of this one to find
	// the closing quote and combine them to get the value
	var nameParts []string

	// Add the current part to the name parts since we detected a quote
	nameParts = append(nameParts, strings.ReplaceAll(part, stopChar, ""))
	for shadowIdx, shadowPart := range parts {
		// Skip over parts we've already gone through
		if shadowIdx <= index {
			continue
		}

		nameParts = append(nameParts, shadowPart)
		// When we've found the quote's friend, break out of this loop
		if strings.Contains(shadowPart, stopChar) {
			break
		}
	}

	value := strings.Join(nameParts, " ")
	criteria[expansion] = value
}
