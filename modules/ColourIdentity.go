package modules

import "strings"

const CWHITE = "w"
const CBLUE = "u"
const CBLACK = "b"
const CRED = "r"
const CGREEN = "g"
const CCOLORLESS = "c"

var colours = map[string]string{
	CWHITE:     "white",
	CBLUE:      "blue",
	CBLACK:     "black",
	CRED:       "red",
	CGREEN:     "green",
	CCOLORLESS: "colorless",
}

func ParseColourIdentity(stopChar string, part string, parts []string) ([]string, error) {
	// Strip the stopChar from the part
	// Split the remaining string into parts
	// For each remaining part
	// Stop at the first match for the part
	// Match part to expansion
	// Skip further processing

	matchedColours := make([]string, 0)
	chunk := strings.Split(part, stopChar)[1]
	chunkParts := strings.Split(chunk, "")

	for _, chunkPart := range chunkParts {
		switch chunkPart {
		case CWHITE:
			matchedColours = append(matchedColours, colours[CWHITE])
			continue
		case CBLUE:
			matchedColours = append(matchedColours, colours[CBLUE])
			continue
		case CBLACK:
			matchedColours = append(matchedColours, colours[CBLACK])
			continue
		case CRED:
			matchedColours = append(matchedColours, colours[CRED])
			continue
		case CGREEN:
			matchedColours = append(matchedColours, colours[CGREEN])
			continue
		}
	}

	return matchedColours, nil
}
