package main

import "fmt"

#define EMITS(type) EMIT(type, ts, te)
#define EMIT(type, ts, te) EMIT_TEXT(type, ts, te, nil)
#define EMIT_TEXT(type, ts, te, text) tokens = append(tokens, Token{ TokenType: TokenType(type), Pos: Pos { ts, te, lineCount+1, (ts-lineStart)+1 }, Text: text })

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

func lex(data []byte) ([]Token, error) {
%% machine scanner;
%% write data;

  // make go stop complaining about unused variables
  _ = scanner_start
  _ = scanner_first_final
  _ = scanner_error
  _ = scanner_en_main
  //_ = _scanner_nfa_targs
  //_ = _scanner_nfa_offsets
  //_ = _scanner_nfa_push_actions
  //_ = _scanner_nfa_pop_trans
  _ = scanner_en_inside_dollar_curly
  _ = scanner_en_string
  _ = scanner_en_ind_string

  cs, p, pe, eof := 0, 0, len(data), len(data)
  ts, te, act := 0, 0, 0

  top := 0
  stack := make([]int, 4, 4)

  var lineStart int
  var lineCount int

  _, _, _, _, _, _ = top, ts, te, act, eof, stack

  // Try to reserve enough capacity so we can avoid copying.
  // This factor is based on lexing all-packages.nix, where we reduce our lexing
  // time to one third compared to the naive empty slice.
  tokens := make([]Token, 0, len(data) / 10)
  //tokens := make([]Token, 0, 0)
  _ = tokens

%%{

  # This machine matches everything, taking note of newlines.
  NL        =  ( any | '\n' @{ lineStart = p+1; lineCount++ } )*;

  ID        =  [a-zA-Z_] [a-zA-Z0-9_'\-]*;
  INT       =  [0-9]+;
  FLOAT     =  (([1-9][0-9]* "." [0-9]*) | (0? "." [0-9]+)) ([Ee][+\-]?[0-9]+)?;
  PATH      =  [a-zA-Z0-9._\-+]* ("/" [a-zA-Z0-9._\-+]+)+;
  HPATH     =  "~" ("/" [a-zA-Z0-9._\-+]+)+;
  SPATH     =  "<" [a-zA-Z0-9._\-+]+ ("/" [a-zA-Z0-9._\-+]+)* ">";
  URI       =  [a-zA-Z] [a-zA-Z0-9+\-.]* ":" [a-zA-Z0-9%/?:@&=+$,\-_.!~*']+;
  SCOMMENT  =  /#[^\r\n]*/;
  LCOMMENT  =  ('/*' ( any* - ( any* '*/' any* ) ) '*/') & NL;

  string := |*
    ( [^$"\\] | '$' [^{"\\] | '\\' (any & NL) | '$\\' (any & NL) )* '$"'
                  { EMIT_TEXT(STR, ts, te-1, data[ts:te-1]);
                    EMIT(DQUOTE, te-1, te);
                    fret;
                  };
    ( [^$"\\] | '$' [^{"\\] | '\\' (any & NL) | '$\\' (any & NL) )+
                  { EMIT_TEXT(STR, ts, te, data[ts:te])
                  };
    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    '"'           { EMITS(DQUOTE);
                    fret; };
  *|;

  ind_string := |*
    ( [^$'] | '$' [^{'] | "'" [^'$] )+ & NL
                        { EMIT_TEXT(IND_STR, ts, te, data[ts:te]) };
    "''$"               { EMIT_TEXT(IND_STR, ts, te, []byte("$")) };
    "'''"               { EMIT_TEXT(IND_STR, ts, te, []byte("'")) };
    "''\\" (any & NL)   { EMIT_TEXT(IND_STR, ts, te, unescapeStr(data[ts+2:te])) };
    "${"                { EMITS(DOLLAR_CURLY);
                          fcall inside_dollar_curly; };
    "''"                { EMITS(IND_STRING_CLOSE);
                          fret; };
    "'"                 { EMIT_TEXT(IND_STR, ts, te, []byte("'")) };
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
    ":"           { EMITS(COLON) };
    ";"           { EMITS(SCOLON) };
    "@"           { EMITS(AT) };
    "-"           { EMITS(MINUS) };
    "+"           { EMITS(PLUS) };
    "!"           { EMITS(NOT) };
    "*"           { EMITS(STAR) };

    "=="          { EMITS(EQ) };
    "!="          { EMITS(NEQ) };
    "<="          { EMITS(LEQ) };
    ">="          { EMITS(GEQ) };
    "&&"          { EMITS(AND) };
    "||"          { EMITS(OR) };
    "->"          { EMITS(IMPL) };
    "//"          { EMITS(UPDATE) };
    "++"          { EMITS(CONCAT) };

    ID            { EMIT_TEXT(ID, ts, te, data[ts:te]) };
    INT           { EMIT_TEXT(INT, ts, te, data[ts:te]) };
    FLOAT         { EMIT_TEXT(FLOAT, ts, te, data[ts:te]) };

    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { EMITS(RCURLY)
                    fret; };
    "{"           { EMITS(LCURLY)
                    fcall inside_dollar_curly; };

    "''" (' '* ('\n' & NL))?
                  { EMITS(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { EMITS(PATH) };
    HPATH         { EMITS(HPATH) };
    SPATH         { EMITS(SPATH) };
    URI           { EMITS(URI) };

    [ \t\r\n]+ & NL;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT) };

    '"'           { EMITS(DQUOTE);
                    fcall string; };

    "="           { EMITS(EQS) };
    "?"           { EMITS(QMARK) };
    ","           { EMITS(COMMA) };
    "("           { EMITS(LPAREN) };
    ")"           { EMITS(RPAREN) };
    "["           { EMITS(LBRACK) };
    "]"           { EMITS(RBRACK) };
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
    ":"           { EMITS(COLON) };
    ";"           { EMITS(SCOLON) };
    "@"           { EMITS(AT) };
    "-"           { EMITS(MINUS) };
    "+"           { EMITS(PLUS) };
    "!"           { EMITS(NOT) };
    "*"           { EMITS(STAR) };

    "=="          { EMITS(EQ) };
    "!="          { EMITS(NEQ) };
    "<="          { EMITS(LEQ) };
    ">="          { EMITS(GEQ) };
    "&&"          { EMITS(AND) };
    "||"          { EMITS(OR) };
    "->"          { EMITS(IMPL) };
    "//"          { EMITS(UPDATE) };
    "++"          { EMITS(CONCAT) };

    ID            { EMIT_TEXT(ID, ts, te, data[ts:te]) };
    INT           { EMIT_TEXT(INT, ts, te, data[ts:te]) };
    FLOAT         { EMIT_TEXT(FLOAT, ts, te, data[ts:te]) };

    "${"          { EMITS(DOLLAR_CURLY);
                    fcall inside_dollar_curly; };
    "}"           { EMITS(RCURLY) };
    "{"           { EMITS(LCURLY) };

    "''" (' '* ('\n' & NL))?
                  { EMITS(IND_STRING_OPEN);
                    fcall ind_string; };

    PATH          { EMITS(PATH) };
    HPATH         { EMITS(HPATH) };
    SPATH         { EMITS(SPATH) };
    URI           { EMITS(URI) };

    [ \t\r\n]+ & NL;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT) };

    '"'           { EMITS(DQUOTE);
                    fcall string; };

    "="           { EMITS(EQS) };
    "?"           { EMITS(QMARK) };
    ","           { EMITS(COMMA) };
    "("           { EMITS(LPAREN) };
    ")"           { EMITS(RPAREN) };
    "["           { EMITS(LBRACK) };
    "]"           { EMITS(RBRACK) };
  *|;

  prepush {
    if len(stack) <= top {
      stack = append(stack, 0)
    }
  }
}%%

%% write init;
%% write exec;

  var err error
	if cs < scanner_first_final {
    err = fmt.Errorf("Unexpected token at ", lineCount+1, (p-lineStart)+1)
  }

  return tokens, err
}
