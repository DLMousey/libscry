package structs

type ParseResult struct {
	ColourIdentity []string `json:"colour_identity,omitempty"`
	Power          int      `json:"power,omitempty"`
	Toughness      int      `json:"toughness,omitempty"`
	Loyalty        int      `json:"loyalty,omitempty"`
	Type           string   `json:"type,omitempty"`
}
