package lexer
import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LET, "let"},
		{IDENT, "x"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type error . shouldbe=%q, coming=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - string  error. shouldbe=%q, coming=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}