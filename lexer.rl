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

  STR
  IND_STR
  IND_STRING_OPEN
  IND_STRING_CLOSE

  // ASCII Chars
  LCURLY = 123
  RCURLY = 125
  DOT    = 46
  SQUOTE = 39
  DQUOTE = 34
)

type Token struct {
  TokenType TokenType
  Start     int
  End       int
  Text      []byte
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
}

func (tok TokenType) String() string {
  s := ""
  if 0 <= tok && tok < TokenType(len(tokens)) {
    s = tokens[tok]
  }
  if s == "" {
    /* s = "token(" + strconv.Itoa(int(tok)) + ")" */
    s = string([]byte { byte(tok) })
  }
  return s
}

#define EMITS(type) EMIT(type, ts, te)
#define EMIT(type, ts, te) EMIT_TEXT(type, ts, te, nil)
#define EMIT_TEXT(type, ts, te, text) tokens = append(tokens, Token{ TokenType: TokenType(type), Start: ts, End: te, Text: text })

func unescapeStr(s []byte) []byte {
  t := []byte{}

  idx := 0
  var c byte
  for ; idx < len(s) ; {
    c = s[idx]
    idx++
    switch c {
    case '\\': {
      c = s[idx]
      idx++
      switch c {
      case 'n': t = append(t, byte('\n'))
      case 'r': t = append(t, byte('\r'))
      case 't': t = append(t, byte('\t'))
      default: t = append(t, byte(c))
      }
    }
    case '\r': {
      t = append(t, byte('\n'))
      if s[idx] == '\n' {
        idx++
      }
    }
    default: t = append(t, byte(c))
    }
  }

  return t
}

func lex(data []byte) []Token {
%% machine scanner;
%% write data;

  // make go stop complaining about unused variables
  _ = scanner_start
  _ = scanner_first_final
  _ = scanner_error
  _ = scanner_en_main
  _ = _scanner_nfa_targs
  _ = _scanner_nfa_offsets
  _ = _scanner_nfa_push_actions
  _ = _scanner_nfa_pop_trans
  _ = scanner_en_inside_dollar_curly
  _ = scanner_en_string
  _ = scanner_en_ind_string

  cs, p, pe, eof := 0, 0, len(data), len(data)
  ts, te, act := 0, 0, 0

  top := 0
  stack := make([]int, 4, 4)

  _, _, _, _, _, _ = top, ts, te, act, eof, stack

  tokens := make([]Token, 0, 0)
  _ = tokens

%%{

  ID        =  [a-zA-Z_] [a-zA-Z0-9_'\-]*;
  INT       =  [0-9]+;
  FLOAT     =  (([1-9][0-9]* "." [0-9]*) | (0? "." [0-9]+)) ([Ee][+\-]?[0-9]+)?;
  PATH      =  [a-zA-Z0-9._\-+]* ("/" [a-zA-Z0-9._\-+]+)+;
  HPATH     =  "~" ("/" [a-zA-Z0-9._\-+]+)+;
  SPATH     =  "<" [a-zA-Z0-9._\-+]+ ("/" [a-zA-Z0-9._\-+]+)* ">";
  URI       =  [a-zA-Z] [a-zA-Z0-9+\-.]* ":" [a-zA-Z0-9%/?:@&=+$,\-_.!~*']+;
  SCOMMENT  =  /#[^\r\n]*/;
  LCOMMENT  =  '/*' ( any* - ( any* '*/' any* ) ) '*/';

  string := |*
    ( [^$"\\] | '$' [^{"\\] | '\\' any | '$\\' any )* '$"'
                  { EMIT_TEXT(STR, ts, te-1, data[ts:te-1]);
                    EMIT(DQUOTE, te-1, te);
                    fret;
                  };
    ( [^$"\\] | '$' [^{"\\] | '\\' any | '$\\' any )+
                  { EMIT_TEXT(STR, ts, te, data[ts:te])
                  };
    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    '"'           { EMITS(DQUOTE);
                    fret; };
  *|;

  ind_string := |*
    ( [^$'] | '$' [^{'] | "'" [^'$] )+
                  { EMIT_TEXT(IND_STR, ts, te, data[ts:te]) };
    "''$"         { EMIT_TEXT(IND_STR, ts, te, []byte("$")) };
    "'''"         { EMIT_TEXT(IND_STR, ts, te, []byte("''")) };
    "''\\" any    { EMIT_TEXT(IND_STR, ts, te, unescapeStr(data[ts+2:te])) };
    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "''"          { EMITS(IND_STRING_CLOSE);
                    fret; };
    "'"           { EMIT_TEXT(IND_STR, ts, te, []byte("'")) };
  *|;

  inside_dollar_curly := |*
    "if"          { EMITS(IF) };
    "then"        { EMITS(THEN) };
    "else"        { EMITS(ELSE) };
    "assert"      { EMITS(ASSERT) };
    "with"        { EMITS(WITH) };
    "let"         { EMITS(LET) };
    "in"          { EMITS(IN) };
    "rec"         { EMITS(REC) };
    "inherit"     { EMITS(INHERIT) };
    "or"          { EMITS(OR_KW) };
    "..."         { EMITS(ELLIPSIS) };

    "=="          { EMITS(EQ) };
    "!="          { EMITS(NEQ) };
    "<="          { EMITS(LEQ) };
    ">="          { EMITS(GEQ) };
    "&&"          { EMITS(AND) };
    "||"          { EMITS(OR) };
    "->"          { EMITS(IMPL) };
    "//"          { EMITS(UPDATE) };
    "++"          { EMITS(CONCAT) };

    ID            { EMITS(ID) };
    INT           { EMITS(INT) };
    FLOAT         { EMITS(FLOAT) };

    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { EMITS(RCURLY)
                    fret; };
    "{"           { EMITS(LCURLY)
                    fcall inside_dollar_curly; };

    "''" (' '* '\n')?
                  { EMITS(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { EMITS(PATH) };
    HPATH         { EMITS(HPATH) };
    SPATH         { EMITS(SPATH) };
    URI           { EMITS(URI) };

    [ \t\r\n]+;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT) };

    '"'           { EMITS(DQUOTE);
                    fcall string; };
  *|;

  main := |*
    "if"          { EMITS(IF) };
    "then"        { EMITS(THEN) };
    "else"        { EMITS(ELSE) };
    "assert"      { EMITS(ASSERT) };
    "with"        { EMITS(WITH) };
    "let"         { EMITS(LET) };
    "in"          { EMITS(IN) };
    "rec"         { EMITS(REC) };
    "inherit"     { EMITS(INHERIT) };
    "or"          { EMITS(OR_KW) };
    "..."         { EMITS(ELLIPSIS) };

    "=="          { EMITS(EQ) };
    "!="          { EMITS(NEQ) };
    "<="          { EMITS(LEQ) };
    ">="          { EMITS(GEQ) };
    "&&"          { EMITS(AND) };
    "||"          { EMITS(OR) };
    "->"          { EMITS(IMPL) };
    "//"          { EMITS(UPDATE) };
    "++"          { EMITS(CONCAT) };

    ID            { EMITS(ID) };
    INT           { EMITS(INT) };
    FLOAT         { EMITS(FLOAT) };

    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { EMITS(RCURLY) };
    "{"           { EMITS(LCURLY) };

    "''" (' '* '\n')?
                  { EMITS(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { EMITS(PATH) };
    HPATH         { EMITS(HPATH) };
    SPATH         { EMITS(SPATH) };
    URI           { EMITS(URI) };

    [ \t\r\n]+;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT) };

    '"'           { EMITS(DQUOTE);
                    fcall string; };
  *|;

  prepush {
    if len(stack) <= top {
      stack = append(stack, 0)
    }
  }
}%%

%% write init;
%% write exec;

  return tokens
}
