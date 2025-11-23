package model

type StandardField struct {
	ID          string
	Key         string // canonical name: "expense.amount"
	Type        string // "string" | "number" | "date" | "datetime" | "array" | "object"
	Description string
	Category    string   // expense | schedule | emotion | ...
	Intents     []string // ["expense"], ["emotion"], ...
	Examples    []any
	Deprecated  bool
	Version     int
	Tags        []string
}
