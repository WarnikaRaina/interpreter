package token

type TokenType string

const(
	ILLEGAL="ILLEGAL"  //special types- a token/char we don't know of 
	EOF="EOF"           //special types- end of file, tells parser to stop

	// Identifiers + literals
	IDENT="IDENT"
	INT="INT"

	// Operators
	ASSIGN="="
	PLUS="+"

	// Delimiters
	COMMA=","
	SEMICOLON=";"

	LPAREN="("
	RPAREN=")"
	LBRACE="{"
	RBRACE="}"

	// Keywords
	FUNCTION="FUNCTION"
	LET="LET"
)

type Token struct{
	type TokenType
	literal string
}