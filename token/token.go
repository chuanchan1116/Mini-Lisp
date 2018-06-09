package token

type TokenType int

const (
	_ = iota
	NUM
	BOOL
	ID
	DEFINE
	FUNC
	IF
	PLUS
	MINUS
	MUL
	DIV
	MOD
	GREATER
	SMALLER
	EQU
	AND
	OR
	NOT
	LPARAM
	RPARAM
	PRINTNUM
	PRINTBOOL
	EOL
)

type Token struct {
	Type TokenType
	Data string
}
