package structs

type ParseResult struct {
	ColourIdentity []string     `json:"colour_identity,omitempty"`
	Power          NumericValue `json:"power,omitempty"`
	Toughness      NumericValue `json:"toughness,omitempty"`
	Loyalty        NumericValue `json:"loyalty,omitempty"`
	Type           string       `json:"type,omitempty"`
	Rarity         []string     `json:"rarity,omitempty"`
	ManaCost       NumericValue `json:"mana_cost,omitempty"`
}
