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
