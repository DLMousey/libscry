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

var ReverseOperators = map[string]string{
	"lt":   CLESSTHAN,
	"lteq": CLESSTHANEQUALTO,
	"eq":   CEQUALS,
	"neq":  CGREATERTHAN,
	"gt":   CGREATERTHANEQUALTO,
	"gteq": CGREATERTHANEQUALTO,
}
