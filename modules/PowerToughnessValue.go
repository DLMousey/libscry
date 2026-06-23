package modules

import (
	"libscry/structs"
	"regexp"
	"strings"
)

const CLESSTHAN = "<"
const CLESSTHANEQUALTO = "<="
const CEQUALS = "="
const CNOTEQUALS = "!="
const CGREATERTHANEQUALTO = ">="
const CGREATERTHAN = ">"

var operators = map[string]string{
	CLESSTHAN:           "lt",
	CLESSTHANEQUALTO:    "lteq",
	CEQUALS:             "eq",
	CNOTEQUALS:          "neq",
	CGREATERTHAN:        "gt",
	CGREATERTHANEQUALTO: "gteq",
}

func parseNumericValue(stopChar string, part string, parts []string) (structs.NumericValue, error) {
	// Strip the stopChar from the part
	// Run a substring match for any operator
	// Find the end index of the operator
	// Split the string between the operator and the value

	// Sensible default of explicit equals operator (pow:5 would be an eq query)
	matchedOperator := CEQUALS

	// Flag to set whether an operator was provided, which changes value extraction behaviour
	hadOperator := false

	comparison := strings.ReplaceAll(part, stopChar, "")
	nonOpPattern := regexp.MustCompile(`([^<=>])`)
	nComparison := nonOpPattern.ReplaceAllString(comparison, "")

	for key, _ := range operators {
		// Will need to manually check this rather than using contains as it does partial matches
		if key == nComparison {
			matchedOperator = key
			hadOperator = true
			break
		}
	}

	var val string

	// Now we have an operator (which may be more than 1 character) we need
	// to remove the stop char and the operator to leave just the value.
	if hadOperator {
		val = strings.ReplaceAll(part, stopChar+matchedOperator, "")
	} else {
		val = strings.ReplaceAll(part, stopChar, "")
	}

	result := structs.NumericValue{
		Operator: operators[matchedOperator],
		Value:    val,
	}

	return result, nil
}

func ParsePowerValue(stopChar string, part string, parts []string) (structs.NumericValue, error) {
	return parseNumericValue(stopChar, part, parts)
}

func ParseToughnessValue(stopChar string, part string, parts []string) (structs.NumericValue, error) {
	return parseNumericValue(stopChar, part, parts)
}

func ParseLoyaltyValue(stopChar string, part string, parts []string) (structs.NumericValue, error) {
	return parseNumericValue(stopChar, part, parts)
}
