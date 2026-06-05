package main

import "testing"

func TestInstantAbbreviation(t *testing.T) {
	query := "t:instant"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "instant"}

	if parts["type"] == "" {
		t.Error("No 'type' part in results")
	}

	if parts["type"] != expectedPart["type"] {
		t.Error("Expected part[0] to be type:instant")
	}
}

func TestInstantFull(t *testing.T) {
	query := "type:instant"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "instant"}

	if parts["type"] == "" {
		t.Error("No 'type' part in results")
	}

	if parts["type"] != expectedPart["type"] {
		t.Error("Expected part[0] to be type:instant")
	}
}

func TestCreatureAbbreviation(t *testing.T) {
	query := "t:creature"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "creature"}

	if parts["type"] == "" {
		t.Error("No 'type' part in results")
	}

	if parts["type"] != expectedPart["type"] {
		t.Error("Expected part[0] to be type:creature")
	}
}

func TestCreatureFull(t *testing.T) {
	query := "type:creature"
	parts := Parse(query)

	expectedPart := map[string]string{"type": "creature"}

	if parts["type"] == "" {
		t.Error("No 'type' part in results")
	}

	if parts["type"] != expectedPart["type"] {
		t.Error("Expected part[0] to be type:creature")
	}
}

func TestColourAbbreviation(t *testing.T) {
	query := "c:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if parts["color"] == "" {
		t.Error("No 'color' part in results")
	}

	if parts["color"] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestColourFull(t *testing.T) {
	query := "color:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if parts["color"] == "" {
		t.Error("No 'color' part in results")
	}

	if parts["color"] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestColourFullCorrect(t *testing.T) {
	query := "colour:g"
	parts := Parse(query)

	expectedPart := map[string]string{"color": "green"}

	if parts["color"] == "" {
		t.Error("No 'color' part in results")
	}

	if parts["color"] != expectedPart["color"] {
		t.Error("Expected part[0] to be color with value 'green'")
	}
}

func TestConsolidated(t *testing.T) {
	query := "pow:5 tou:5 loy:5 type:creature"
	parts := Parse(query)

	expectedParts := map[string]string{
		"power":     "5",
		"toughness": "5",
		"loyalty":   "5",
		"type":      "creature",
	}

	hasPower := parts["power"] != ""
	hasToughness := parts["toughness"] != ""
	hasLoyalty := parts["loyalty"] != ""
	hasType := parts["type"] != ""

	matchedPower := parts["power"] == expectedParts["power"]
	matchedToughness := parts["toughness"] == expectedParts["toughness"]
	matchedLoyalty := parts["loyalty"] == expectedParts["loyalty"]
	matchedType := parts["type"] == expectedParts["type"]

	if !hasPower || !hasToughness || !hasLoyalty || !hasType {
		t.Error("Missing required part from parsed tokens")
	}

	if !matchedPower || !matchedToughness || !matchedLoyalty || !matchedType {
		t.Error("Part mismatch")
	}
}
