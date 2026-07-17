package main

import (
	"libscry/modules"
	"libscry/structs"
	"strings"
)

const TDOUBLEQUOTE = "\""
const TSINGLEQUOTE = "'"
const TSPACE = " "
const TOR = "or"
const TAND = "and"
const TTYPE = "t:"
const TTYPEFULL = "type:"
const TCOLOR = "c:"
const TCOLORFULL = "color:"
const TCOLORCORRECTFULL = "colour:"
const TORACLE = "o:"
const TORACLEFULL = "oracle:"
const TPOWERFULL = "power:"
const TPOWER = "pow:"
const TTOUGHNESSFULL = "toughness:"
const TTOUGHNESS = "tou:"
const TLOYALTYFULL = "loyalty:"
const TLOYALTY = "loy:"
const TMANACOST = "m:"
const TMANACOSTFULL = "mana:"
const TRARITY = "r:"
const TRARITYFULL = "rarity:"

const ETYPE = "type"
const ENAME = "name"
const ECOLOR = "color"
const EORACLE = "oracle"
const EPOWER = "power"
const ETOUGHNESS = "toughness"
const ELOYALTY = "loyalty"
const EMANA = "mana"
const ERARITY = "rarity"

var expansions = map[string]string{
	TTYPE:             ETYPE,
	TTYPEFULL:         ETYPE,
	TDOUBLEQUOTE:      ENAME,
	TSINGLEQUOTE:      ENAME,
	TCOLOR:            ECOLOR,
	TCOLORFULL:        ECOLOR,
	TCOLORCORRECTFULL: ECOLOR,
	TORACLE:           EORACLE,
	TORACLEFULL:       EORACLE,
	TPOWERFULL:        EPOWER,
	TPOWER:            EPOWER,
	TTOUGHNESSFULL:    ETOUGHNESS,
	TTOUGHNESS:        ETOUGHNESS,
	TLOYALTYFULL:      ELOYALTY,
	TLOYALTY:          ELOYALTY,
	TMANACOST:         EMANA,
	TMANACOSTFULL:     EMANA,
	TRARITY:           ERARITY,
	TRARITYFULL:       ERARITY,
}

var stopChars []string = []string{
	TSPACE, TOR, TAND, TTYPE, TTYPEFULL, TCOLOR,
	TCOLORFULL, TCOLORCORRECTFULL, TORACLE,
	TORACLEFULL, TDOUBLEQUOTE, TSINGLEQUOTE,
	TPOWER, TPOWERFULL, TTOUGHNESS, TTOUGHNESSFULL,
	TLOYALTY, TLOYALTYFULL, TMANACOST, TMANACOSTFULL,
	TRARITY, TRARITYFULL,
}

func Parse(input string) structs.ParseResult {
	// Break string into parts on each space
	// Iterate through each part
	// If a part begins with a quote mark, find the next part that contains another quote to delimit it,
	// then join these parts into a single part and remove the separate ones from the array

	// Some sanitisation first - replace all single quotes with double quotes,
	// strip the newline suffix and lowercase the whole thing
	input = strings.ReplaceAll(input, "'", "\"")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ToLower(input)
	parts := strings.Split(input, TSPACE)

	structCriteria := structs.ParseResult{
		ColourIdentity: []string{},
	}

	for _, part := range parts {
		// Iterate over each of the tokens for each part and see if it's contained within this part
		for _, stopChar := range stopChars {
			hastoken := strings.Contains(part, stopChar)

			if hastoken {

				// find the index
				startIndex := strings.Index(part, stopChar)
				endIndex := startIndex + len(stopChar)

				marker := part[startIndex:endIndex]
				expansion := expansions[marker]

				switch expansion {
				case ECOLOR:
					colours, err := modules.ParseColourIdentity(stopChar, part, parts)
					structCriteria.ColourIdentity = append(structCriteria.ColourIdentity, colours...)

					if err != nil {
						panic(err)
					}
					break
				case EPOWER:
					structCriteria.Power, _ = modules.ParsePowerValue(stopChar, part, parts)
					break
				case ETOUGHNESS:
					structCriteria.Toughness, _ = modules.ParseToughnessValue(stopChar, part, parts)
					break
				case ELOYALTY:
					structCriteria.Loyalty, _ = modules.ParseLoyaltyValue(stopChar, part, parts)
					break
				case ETYPE:
					structCriteria.Type = part[endIndex:]
					break
				case ERARITY:
					structCriteria.Rarity, _ = modules.ParseRarity(stopChar, part, parts)
					break
				case EMANA:
					structCriteria.ManaCost, _ = modules.ParseManaValue(stopChar, part, parts)
					break
				}
			}
		}
	}

	return structCriteria
}

func extractName(criteria map[string]string, index int, stopChar string, part string, parts []string, expansion string) {
	// First of all check if the current part contains 2 quotes.
	// If there is 2 we'll skip the lookahead and add it straight into the criteria
	quoteCount := strings.Count(part, TDOUBLEQUOTE)
	if quoteCount == 2 {
		criteria[expansion] = part
		return
	}

	// Otherwise we'll have to look through each part in front of this one to find
	// the closing quote and combine them to get the value
	var nameParts []string

	// Add the current part to the name parts since we detected a quote
	nameParts = append(nameParts, strings.ReplaceAll(part, stopChar, ""))
	for shadowIdx, shadowPart := range parts {
		// Skip over parts we've already gone through
		if shadowIdx <= index {
			continue
		}

		nameParts = append(nameParts, shadowPart)
		// When we've found the quote's friend, break out of this loop
		if strings.Contains(shadowPart, stopChar) {
			break
		}
	}

	value := strings.Join(nameParts, " ")
	criteria[expansion] = value
}
