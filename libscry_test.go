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
