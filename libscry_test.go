package main

import (
	"libscry/structs"
	"strings"
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

func TestPowerNotEqualTo(t *testing.T) {
	query := "pow:!=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "neq",
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

func TestToughnessImplicit(t *testing.T) {
	query := "tou:5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "eq",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestToughnessNotEqualTo(t *testing.T) {
	query := "tou:!=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "neq",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestToughnessLessThan(t *testing.T) {
	query := "tou:<5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lt",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestToughnessLessThanEqualTo(t *testing.T) {
	query := "tou:<=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lteq",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestToughnessGreaterThan(t *testing.T) {
	query := "tou:>5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gt",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestToughnessGreaterThanEqualTo(t *testing.T) {
	query := "tou:>=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gteq",
		Value:    "5",
	}

	if parts.Toughness.Operator != expectedPart.Operator {
		t.Error("Expected toughness operator to be " + expectedPart.Operator + ", got " + parts.Toughness.Operator)
	}

	if parts.Toughness.Value != expectedPart.Value {
		t.Error("Expected toughness value to be " + expectedPart.Value + ", got " + parts.Toughness.Value)
	}
}

func TestLoyaltyImplicit(t *testing.T) {
	query := "loy:5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "eq",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestLoyaltyNotEqualTo(t *testing.T) {
	query := "loy:!=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "neq",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestLoyaltyLessThan(t *testing.T) {
	query := "loy:<5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lt",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestLoyaltyLessThanEqualTo(t *testing.T) {
	query := "loy:<=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "lteq",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestLoyaltyGreaterThan(t *testing.T) {
	query := "loy:>5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gt",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestLoyaltyGreaterThanEqualTo(t *testing.T) {
	query := "loy:>=5"
	parts := Parse(query)

	expectedPart := structs.NumericValue{
		Operator: "gteq",
		Value:    "5",
	}

	if parts.Loyalty.Operator != expectedPart.Operator {
		t.Error("Expected loyalty operator to be " + expectedPart.Operator + ", got " + parts.Loyalty.Operator)
	}

	if parts.Loyalty.Value != expectedPart.Value {
		t.Error("Expected loyalty value to be " + expectedPart.Value + ", got " + parts.Loyalty.Value)
	}
}

func TestRarityMythicSingle(t *testing.T) {
	query := "r:m"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "mythic")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity to be " + expectedPart[0] + ", got " + strings.Join(parts.Rarity, ", "))
	}
}

func TestRarityRareSingle(t *testing.T) {
	query := "r:r"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "rare")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity to be " + expectedPart[0] + ", got " + strings.Join(parts.Rarity, ", "))
	}
}

func TestRarityUncommonSingle(t *testing.T) {
	query := "r:u"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "uncommon")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity to be " + expectedPart[0] + ", got " + strings.Join(parts.Rarity, ", "))
	}
}

func TestRarityCommonSingle(t *testing.T) {
	query := "r:c"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "common")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity to be " + expectedPart[0] + ", got " + strings.Join(parts.Rarity, ", "))
	}
}

func TestRarityMythicRareMultiple(t *testing.T) {
	query := "r:m,r"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "mythic")
	expectedPart = append(expectedPart, "rare")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity[0] to be " + expectedPart[0] + ", got " + parts.Rarity[0])
	}

	if parts.Rarity[1] != expectedPart[1] {
		t.Error("Expected rarit[1] to be " + expectedPart[1] + ", got " + parts.Rarity[1])
	}
}

func TestRarityUncommonCommonMultiple(t *testing.T) {
	query := "r:u,c"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "uncommon")
	expectedPart = append(expectedPart, "common")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity[0] to be " + expectedPart[0] + ", got " + parts.Rarity[0])
	}

	if parts.Rarity[1] != expectedPart[1] {
		t.Error("Expected rarit[1] to be " + expectedPart[1] + ", got " + parts.Rarity[1])
	}
}

func TestRaritySpecialBonusMultiple(t *testing.T) {
	query := "r:s,b"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "special")
	expectedPart = append(expectedPart, "bonus")

	if parts.Rarity[0] != expectedPart[0] {
		t.Error("Expected rarity[0] to be " + expectedPart[0] + ", got " + parts.Rarity[0])
	}

	if parts.Rarity[1] != expectedPart[1] {
		t.Error("Expected rarity[1] to be " + expectedPart[1] + ", got " + parts.Rarity[1])
	}
}

func TestRarityAllValuesMultiple(t *testing.T) {
	query := "r:m,r,u,c,s,b"
	parts := Parse(query)

	expectedPart := make([]string, 0)
	expectedPart = append(expectedPart, "mythic")
	expectedPart = append(expectedPart, "rare")
	expectedPart = append(expectedPart, "uncommon")
	expectedPart = append(expectedPart, "common")
	expectedPart = append(expectedPart, "special")
	expectedPart = append(expectedPart, "bonus")

	for idx, part := range parts.Rarity {
		if part != expectedPart[idx] {
			t.Error("Expected rarity[" + string(rune(idx)) + "] to be " + expectedPart[idx] + ", got " + part)
		}
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
