package main

import (
  "log"
  /* "strconv" */
)

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

type token struct {
  token int
  start int
  end   int
}

type Token int

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

  IND_STR:          "IND_STR",
  IND_STRING_OPEN:  "IND_STRING_OPEN",
  IND_STRING_CLOSE: "IND_STRING_CLOSE",
}

func (tok Token) String() string {
  s := ""
  if 0 <= tok && tok < Token(len(tokens)) {
    s = tokens[tok]
  }
  if s == "" {
    /* s = "token(" + strconv.Itoa(int(tok)) + ")" */
    s = string([]byte { byte(tok) })
  }
  return s
}

#define TOK(type) log.Println("TOK:", Token(type), ts, te)

func lex(data []byte) {
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

  tokens := make([]token, 0, 0)
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
                  { log.Println("STR:", string(data[ts:(te-1)]));
                    TOK(DQUOTE);
                    fret;
                  };
    ( [^$"\\] | '$' [^{"\\] | '\\' any | '$\\' any )+
                  { log.Println("STR:", string(data[ts:(te)]));
                  };
    "${"          { TOK(DOLLAR_CURLY); fcall inside_dollar_curly; };
    '"'           { TOK(DQUOTE); fret; };
  *|;

  ind_string := |*
    ( [^$'] | '$' [^{'] | "'" [^'$] )+
                  { TOK(IND_STR) };
    "''$"         { TOK(IND_STR) };
    "'''"         { TOK(IND_STR) };
    "''\\" any    { TOK(IND_STR) };
    "${"          { TOK(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "''"          { TOK(IND_STRING_CLOSE);
                    fret; };
    "'"           { TOK(IND_STR) };
  *|;

  inside_dollar_curly := |*
    "if"          { TOK(IF) };
    "then"        { TOK(THEN) };
    "else"        { TOK(ELSE) };
    "assert"      { TOK(ASSERT) };
    "with"        { TOK(WITH) };
    "let"         { TOK(LET) };
    "in"          { TOK(IN) };
    "rec"         { TOK(REC) };
    "inherit"     { TOK(INHERIT) };
    "or"          { TOK(OR_KW) };
    "..."         { TOK(ELLIPSIS) };

    "=="          { TOK(EQ) };
    "!="          { TOK(NEQ) };
    "<="          { TOK(LEQ) };
    ">="          { TOK(GEQ) };
    "&&"          { TOK(AND) };
    "||"          { TOK(OR) };
    "->"          { TOK(IMPL) };
    "//"          { TOK(UPDATE) };
    "++"          { TOK(CONCAT) };

    ID            { TOK(ID) };
    INT           { TOK(INT) };
    FLOAT         { TOK(FLOAT) };

    "${"          { TOK(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { TOK(RCURLY)
                    fret; };
    "{"           { TOK(LCURLY)
                    fcall inside_dollar_curly; };

    "''" (' '* '\n')?
                  { TOK(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { TOK(PATH) };
    HPATH         { TOK(HPATH) };
    SPATH         { TOK(SPATH) };
    URI           { TOK(URI) };

    [ \t\r\n]+;

    SCOMMENT;
    LCOMMENT;

    "."           { TOK(DOT) };

    '"'           { TOK(DQUOTE);
                    fcall string; };
  *|;

  main := |*
    "if"          { TOK(IF) };
    "then"        { TOK(THEN) };
    "else"        { TOK(ELSE) };
    "assert"      { TOK(ASSERT) };
    "with"        { TOK(WITH) };
    "let"         { TOK(LET) };
    "in"          { TOK(IN) };
    "rec"         { TOK(REC) };
    "inherit"     { TOK(INHERIT) };
    "or"          { TOK(OR_KW) };
    "..."         { TOK(ELLIPSIS) };

    "=="          { TOK(EQ) };
    "!="          { TOK(NEQ) };
    "<="          { TOK(LEQ) };
    ">="          { TOK(GEQ) };
    "&&"          { TOK(AND) };
    "||"          { TOK(OR) };
    "->"          { TOK(IMPL) };
    "//"          { TOK(UPDATE) };
    "++"          { TOK(CONCAT) };

    ID            { TOK(ID) };
    INT           { TOK(INT) };
    FLOAT         { TOK(FLOAT) };

    "${"          { TOK(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { TOK(RCURLY) };
    "{"           { TOK(LCURLY) };

    "''" (' '* '\n')?
                  { TOK(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { TOK(PATH) };
    HPATH         { TOK(HPATH) };
    SPATH         { TOK(SPATH) };
    URI           { TOK(URI) };

    [ \t\r\n]+;

    SCOMMENT;
    LCOMMENT;

    "."           { TOK(DOT) };

    '"'           { TOK(DQUOTE);
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

}
