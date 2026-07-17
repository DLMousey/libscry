package main

import (
	"libscry/structs"
	"testing"
)

func CompareManaValueResult(parts structs.ParseResult, expectedPart structs.NumericValue, t *testing.T) {
	if parts.ManaCost != expectedPart {
		if parts.ManaCost.Operator != expectedPart.Operator {
			t.Error("Expected mana cost operator to be " + expectedPart.Operator + ", got " + parts.ManaCost.Operator)
		}

		if parts.ManaCost.Value != expectedPart.Value {
			t.Error("Expected mana cost value to be " + expectedPart.Value + ", got " + parts.ManaCost.Value)
		}
	}
}
