package lexer


type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
// Identifiers + Literals
	IDENT = "IDENT"

	INT = "INT"
	// operators
	ASSIGN = "="
	PLUS = "+"
	REMOVAL = "-"
	MULTIPLICATION = "*"
	DIVISION = "/"
	//delimnters 
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	LET = "LET"
	PRINTLN = "PRINTLN"
	SCANLN = "SCANLN"

)