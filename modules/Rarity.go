package modules

import "strings"

const CMYTHIC = "m"
const CRARE = "r"
const CUNCOMMON = "u"
const CCOMMON = "c"
const CSPECIAL = "s"
const CBONUS = "b"

const CMYTHICFULL = "mythic"
const CRAREFULL = "rare"
const CUNCOMMONFULL = "uncommon"
const CCOMMONFULL = "common"
const CSPECIALFULL = "special"
const CBONUSFULL = "bonus"

var rarities = map[string]string{
	CMYTHIC:   "mythic",
	CRARE:     "rare",
	CUNCOMMON: "uncommon",
	CCOMMON:   "common",
	CSPECIAL:  "special",
	CBONUS:    "bonus",
}

func ParseRarity(stopChar string, part string, parts []string) ([]string, error) {
	matchedRarities := make([]string, 0)
	chunk := strings.Split(part, stopChar)[1]
	chunkParts := strings.Split(chunk, "")

	for _, chunkPart := range chunkParts {
		switch chunkPart {
		case CMYTHIC, CMYTHICFULL:
			matchedRarities = append(matchedRarities, rarities[CMYTHIC])
			continue
		case CRARE, CRAREFULL:
			matchedRarities = append(matchedRarities, rarities[CRARE])
			continue
		case CUNCOMMON, CUNCOMMONFULL:
			matchedRarities = append(matchedRarities, rarities[CUNCOMMON])
			continue
		case CCOMMON, CCOMMONFULL:
			matchedRarities = append(matchedRarities, rarities[CCOMMON])
			continue
		case CSPECIAL, CSPECIALFULL:
			matchedRarities = append(matchedRarities, rarities[CSPECIAL])
			continue
		case CBONUS, CBONUSFULL:
			matchedRarities = append(matchedRarities, rarities[CBONUS])
			continue
		}
	}

	return matchedRarities, nil
}
