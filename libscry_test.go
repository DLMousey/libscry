package main

import (
	"libscry/structs"
	"testing"
)

func TestInstantAbbreviation(t *testing.T) {
	query := "t:instant"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "instant"}

	if parts.Type == "" {
		t.Error("No 'type' part in results")
	}

	if parts.Type != expectedPart["type"] {
		t.Error("Expected part[0] to be type:instant")
	}
}

func TestInstantFull(t *testing.T) {
	query := "type:instant"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "instant"}

	if parts.Type == "" {
		t.Error("No 'type' part in results")
	}

	if parts.Type != expectedPart["type"] {
		t.Error("Expected part[0] to be type:instant")
	}
}

func TestCreatureAbbreviation(t *testing.T) {
	query := "t:creature"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "creature"}

	if parts.Type == "" {
		t.Error("No 'type' part in results")
	}

	if parts.Type != expectedPart["type"] {
		t.Error("Expected part[0] to be type:creature")
	}
}

func TestCreatureFull(t *testing.T) {
	query := "type:creature"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "creature"}

	if parts.Type == "" {
		t.Error("No 'type' part in results")
	}

	if parts.Type != expectedPart["type"] {
		t.Error("Expected part[0] to be type:creature")
	}
}

func TestColourAbbreviation(t *testing.T) {
	query := "c:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if len(parts.ColourIdentity) != len(expectedPart) {
		t.Error("No 'color' part in results")
	}

	if parts.ColourIdentity[0] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestColourFull(t *testing.T) {
	query := "color:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if len(parts.ColourIdentity) != len(expectedPart) {
		t.Error("No 'color' part in results")
	}

	if parts.ColourIdentity[0] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestColourFullCorrect(t *testing.T) {
	query := "colour:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if len(parts.ColourIdentity) != len(expectedPart) {
		t.Error("No 'color' part in results")
	}

	if parts.ColourIdentity[0] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestColourMultiple(t *testing.T) {
	query := "c:gr"
	parts := Parse(query)

	expectedColours := []string{"green", "red"}

	resultLength := len(parts.ColourIdentity)
	if resultLength == 0 {
		t.Error("No 'colour' part in results")
	}

	if resultLength != len(expectedColours) {
		t.Error("Missing expected colours from parser results")
	}

	if parts.ColourIdentity[0] != expectedColours[0] {
		t.Error("Expected part[0] to be colour with value 'green'")
	}

	if parts.ColourIdentity[1] != expectedColours[1] {
		t.Error("Expected part[1] to be colour with value 'red'")
	}
}

func TestPowerImplicit(t *testing.T) {
	query := "pow:5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "eq",
		Value:    "5",
	}

	if parts.Power.Operator != expectedPart.Operator {
		t.Error("Expected power operator to be " + expectedPart.Operator + ", got " + parts.Power.Operator)
	}

	if parts.Power.Value != expectedPart.Value {
		t.Error("Expected power value to be " + expectedPart.Value + ", got " + parts.Power.Value)
	}
}

func TestPowerLessThan(t *testing.T) {
	query := "pow:<5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lt",
		Value:    "5",
	}

	if parts.Power.Operator != expectedPart.Operator {
		t.Error("Expected power operator to be " + expectedPart.Operator + ", got " + parts.Power.Operator)
	}

	if parts.Power.Value != expectedPart.Value {
		t.Error("Expected power value to be " + expectedPart.Value + ", got " + parts.Power.Value)
	}
}

func TestPowerLessThanEqualTo(t *testing.T) {
	query := "pow:<=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lteq",
		Value:    "5",
	}

	if parts.Power.Operator != expectedPart.Operator {
		t.Error("Expected power operator to be " + expectedPart.Operator + ", got " + parts.Power.Operator)
	}

	if parts.Power.Value != expectedPart.Value {
		t.Error("Expected power value to be " + expectedPart.Value + ", got " + parts.Power.Value)
	}
}

func TestPowerGreaterThan(t *testing.T) {
	query := "pow:>5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gt",
		Value:    "5",
	}

	if parts.Power.Operator != expectedPart.Operator {
		t.Error("Expected power operator to be " + expectedPart.Operator + ", got " + parts.Power.Operator)
	}

	if parts.Power.Value != expectedPart.Value {
		t.Error("Expected power value to be " + expectedPart.Value + ", got " + parts.Power.Value)
	}
}

func TestPowerGreaterThanEqualTo(t *testing.T) {
	query := "pow:>=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gteq",
		Value:    "5",
	}

	if parts.Power.Operator != expectedPart.Operator {
		t.Error("Expected power operator to be " + expectedPart.Operator + ", got " + parts.Power.Operator)
	}

	if parts.Power.Value != expectedPart.Value {
		t.Error("Expected power value to be " + expectedPart.Value + ", got " + parts.Power.Value)
	}
}

// Disabled due to mana cost requiring rework
//func TestManaCost(t *testing.T) {
//	query := "m:3"
//	parts := Parse(query)
//
//	expectedPart := map[string]string{"mana": "3"}
//
//	if parts["mana"] == "" {
//		t.Error("No 'mana' part in results")
//	}
//
//	if parts["mana"] != expectedPart["mana"] {
//		t.Error("Expected part[0] to be mana with value '3'")
//	}
//}
//
//func TestManaCostFull(t *testing.T) {
//	query := "mana:3"
//	parts := Parse(query)
//
//	expectedPart := map[string]string{"mana": "3"}
//
//	if parts["mana"] == "" {
//		t.Error("No 'mana' part in results")
//	}
//
//	if parts["mana"] != expectedPart["mana"] {
//		t.Error("Expected part[0] to be mana with value '3'")
//	}
//}

//func TestConsolidated(t *testing.T) {
//	query := "pow:5 tou:5 loy:5 type:creature m:3"
//	parts := Parse(query)
//
//	expectedParts := map[string]int{
//		"power":     5,
//		"toughness": 5,
//		"loyalty":   5,
//		//"type":      "creature",
//		//"mana": 3,
//	}
//
//	expectedType := "creature"
//
//	hasPower := parts.Power != {}
//	hasToughness := parts.Toughness != 0
//	hasLoyalty := parts.Loyalty != 0
//	hasType := parts.Type != ""
//
//	//hasMana := parts["mana"] != ""
//
//	matchedPower := parts.Power == expectedParts["power"]
//	matchedToughness := parts.Toughness == expectedParts["toughness"]
//	matchedLoyalty := parts.Loyalty == expectedParts["loyalty"]
//	matchedType := parts.Type == expectedType
//	//matchedMana := parts["mana"] == expectedParts["mana"]
//
//	if !hasPower || !hasToughness || !hasLoyalty || !hasType {
//		t.Error("Missing required part from parsed tokens")
//	}
//
//	if !matchedPower || !matchedToughness || !matchedLoyalty || !matchedType {
//		t.Error("Part mismatch")
//	}
//}
