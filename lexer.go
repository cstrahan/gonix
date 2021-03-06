package main

const (
	IF = iota
	THEN
	ELSE
	ASSERT
	WITH
	LET
	IN
	REC
	INHERIT
	OR_KW
	ELLIPSIS
	DOLLAR_CURLY
	ASSIGN

	EQ
	NEQ
	LEQ
	GEQ
	AND
	OR
	IMPL
	UPDATE
	CONCAT

	ID
	INT
	FLOAT
	PATH
	HPATH
	SPATH
	URI

	// ASCII Chars
	LCURLY = 123
	RCURLY = 125
	DOT    = 46
	SQUOTE = 39
	DQUOTE = 34
	COLON  = 58
	SCOLON = 59
	AT     = 64
	MINUS  = 45
	PLUS   = 43
	NOT    = 33
	QMARK  = 63
	COMMA  = 44
	EQS    = 61
	STAR   = 42
	LPAREN = 40
	RPAREN = 41
	LBRACK = 91
	RBRACK = 93

	// Non-ASCII cont.
	NEGATE           = 126
	STR              = 127
	IND_STR          = 128
	IND_STRING_OPEN  = 129
	IND_STRING_CLOSE = 130
	EOF              = 131
)

type Token struct {
	TokenType TokenType
	Pos
	Text []byte
}

type Pos struct {
	Start int
	End   int
	Row   int
	Col   int
}

type TokenType int

var tokens = [...]string{
	IF:           "IF",
	THEN:         "THEN",
	ELSE:         "ELSE",
	ASSERT:       "ASSERT",
	WITH:         "WITH",
	LET:          "LET",
	IN:           "IN",
	REC:          "REC",
	INHERIT:      "INHERIT",
	OR_KW:        "OR_KW",
	ELLIPSIS:     "ELLIPSIS",
	DOLLAR_CURLY: "DOLLAR_CURLY",

	EQ:     "EQ",
	NEQ:    "NEQ",
	LEQ:    "LEQ",
	GEQ:    "GEQ",
	AND:    "AND",
	OR:     "OR",
	IMPL:   "IMPL",
	UPDATE: "UPDATE",
	CONCAT: "CONCAT",

	ID:    "ID",
	INT:   "INT",
	FLOAT: "FLOAT",
	PATH:  "PATH",
	HPATH: "HPATH",
	SPATH: "SPATH",
	URI:   "URI",

	STR:              "STR",
	IND_STR:          "IND_STR",
	IND_STRING_OPEN:  "IND_STRING_OPEN",
	IND_STRING_CLOSE: "IND_STRING_CLOSE",
	EOF:              "EOF",
}

func (tok TokenType) String() string {
	s := ""
	if 0 <= tok && tok < TokenType(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		/* s = "token(" + strconv.Itoa(int(tok)) + ")" */
		s = string([]byte{byte(tok)})
	}
	return s
}
