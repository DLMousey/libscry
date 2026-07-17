package structs

const CLESSTHAN = "<"
const CLESSTHANEQUALTO = "<="
const CEQUALS = "="
const CNOTEQUALS = "!="
const CGREATERTHANEQUALTO = ">="
const CGREATERTHAN = ">"

var Operators = map[string]string{
	CLESSTHAN:           "lt",
	CLESSTHANEQUALTO:    "lteq",
	CEQUALS:             "eq",
	CNOTEQUALS:          "neq",
	CGREATERTHAN:        "gt",
	CGREATERTHANEQUALTO: "gteq",
}
