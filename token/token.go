package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 식별자 + 리터럴
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// 연산자
	ASSIGN    = "="
	MINUS     = "-"
	PLUS      = "+"
	BANG      = "!"
	ASTERRISK = "*"
	SLASH     = "/"
	EQ        = "=="
	NOT_EQ    = "!="

	LT = "<"
	LE = "<="
	GT = ">"
	GE = ">="

	TILDE        = "~"
	LEFT_SHIFT   = "<<"
	RIGHT_SHIFT  = ">>"
	AMPERSAND    = "&"
	VERTICAL_BAR = "|"
	CARET        = "^"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	CONTINUE = "CONTINUE"
	BREAK    = "BREAK"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"for":      FOR,
	"continue": CONTINUE,
	"break":    BREAK,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
