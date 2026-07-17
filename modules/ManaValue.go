package modules

import (
	"libscry/structs"
	"regexp"
	"strings"
)

func ParseManaValue(stopChar string, part string, parts []string) (structs.NumericValue, error) {
	matchedOperator := structs.CEQUALS

	hadOperator := false

	comparison := strings.ReplaceAll(part, stopChar, "")
	nonOpPattern := regexp.MustCompile(`([^!=><])`)
	nComparison := nonOpPattern.ReplaceAllString(comparison, "")

	for key, _ := range structs.Operators {
		if key == nComparison {
			matchedOperator = key
			hadOperator = true
			break
		}
	}

	var val string

	if hadOperator {
		val = strings.ReplaceAll(part, stopChar+matchedOperator, "")
	} else {
		val = strings.ReplaceAll(part, stopChar, "")
	}

	result := structs.NumericValue{
		Operator: structs.Operators[matchedOperator],
		Value:    val,
	}

	return result, nil
}
