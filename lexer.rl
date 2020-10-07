%% machine scanner;

package main

import "fmt"

%% write data;

#define EMITS(type) EMIT(type, ts, te)
#define EMIT(type, ts, te) EMIT_TEXT(type, ts, te, nil)
#define EMIT_TEXT(type, ts, te, text) token = Token{ TokenType: TokenType(type), Pos: Pos { ts, te, lineCount+1, (ts-lineStart)+1 }, Text: text }

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

func (l *Lexer) Lex() (Token, error) {

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

  // load state from l
  data := l.data
  pe := l.pe
  cs := l.cs
  p := l.p
  top := l.top
  stack := l.stack
  lineStart := l.lineStart
  lineCount := l.lineCount
  act := l.act

  eof := pe
  ts, te, act := 0, 0, 0

  _, _, _, _, _, _ = top, ts, te, act, eof, stack

  // Try to reserve enough capacity so we can avoid copying.
  // This factor is based on lexing all-packages.nix, where we reduce our lexing
  // time to one third compared to the naive empty slice.
  token := Token{}
  _ = token

%% write exec;

  // store state in l
  l.cs = cs
  l.p = p
  l.top = top
  l.lineStart = lineStart
  l.lineCount = lineCount
  l.act = act

  var err error

  if cs == %%{ write error; }%% {
    // TODO: handle error
    err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
  }

	//if cs < scanner_first_final {
  //  err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
  //}

  return token, err
}

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
                    top--
                    fnext *stack[top];
                    fbreak;
                  };
    ( [^$"\\] | '$' [^{"\\] | '\\' (any & NL) | '$\\' (any & NL) )+
                  { EMIT_TEXT(STR, ts, te, data[ts:te]); // ?
                    fbreak; };
    "${"          { EMITS(DOLLAR_CURLY);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext inside_dollar_curly;
                    fbreak;
                    };
    '"'           { EMITS(DQUOTE);
                    top--
                    fnext *stack[top];
                    fbreak;

                    };
  *|;

  ind_string := |*
    ( [^$'] | '$' [^{'] | "'" [^'$] )+ & NL
                        { EMIT_TEXT(IND_STR, ts, te, data[ts:te])
                          fbreak; };
    "''$"               { EMIT_TEXT(IND_STR, ts, te, []byte("$"))
                          fbreak; };
    "'''"               { EMIT_TEXT(IND_STR, ts, te, []byte("'"))
                          fbreak; };
    "''\\" (any & NL)   { EMIT_TEXT(IND_STR, ts, te, unescapeStr(data[ts+2:te]))
                          fbreak; };
    "${"                { EMITS(DOLLAR_CURLY);
                          if len(stack) <= top {
                            stack = append(stack, 0)
                          }
                          stack[top] = ftargs
                          top++
                          fnext inside_dollar_curly;
                          fbreak; };
    "''"                { EMITS(IND_STRING_CLOSE);
                          top--
                          fnext *stack[top];
                          fbreak; };
    "'"                 { EMIT_TEXT(IND_STR, ts, te, []byte("'"))
                          fbreak; };
  *|;

  inside_dollar_curly := |*
    "if"          { EMITS(IF)
                    fbreak; };
    "then"        { EMITS(THEN)
                    fbreak; };
    "else"        { EMITS(ELSE)
                    fbreak; };
    "assert"      { EMITS(ASSERT)
                    fbreak; };
    "with"        { EMITS(WITH)
                    fbreak; };
    "let"         { EMITS(LET)
                    fbreak; };
    "in"          { EMITS(IN)
                    fbreak; };
    "rec"         { EMITS(REC)
                    fbreak; };
    "inherit"     { EMITS(INHERIT)
                    fbreak; };
    "or"          { EMITS(OR_KW)
                    fbreak; };
    "..."         { EMITS(ELLIPSIS)
                    fbreak; };
    ":"           { EMITS(COLON)
                    fbreak; };
    ";"           { EMITS(SCOLON)
                    fbreak; };
    "@"           { EMITS(AT)
                    fbreak; };
    "-"           { EMITS(MINUS)
                    fbreak; };
    "+"           { EMITS(PLUS)
                    fbreak; };
    "!"           { EMITS(NOT)
                    fbreak; };
    "*"           { EMITS(STAR)
                    fbreak; };

    "=="          { EMITS(EQ)
                    fbreak; };
    "!="          { EMITS(NEQ)
                    fbreak; };
    "<="          { EMITS(LEQ)
                    fbreak; };
    ">="          { EMITS(GEQ)
                    fbreak; };
    "&&"          { EMITS(AND)
                    fbreak; };
    "||"          { EMITS(OR)
                    fbreak; };
    "->"          { EMITS(IMPL)
                    fbreak; };
    "//"          { EMITS(UPDATE)
                    fbreak; };
    "++"          { EMITS(CONCAT)
                    fbreak; };


    ID            { EMIT_TEXT(ID, ts, te, data[ts:te])
                    fbreak; };
    INT           { EMIT_TEXT(INT, ts, te, data[ts:te])
                    fbreak; };
    FLOAT         { EMIT_TEXT(FLOAT, ts, te, data[ts:te])
                    fbreak; };

    "${"          { EMITS(DOLLAR_CURLY);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext inside_dollar_curly;
                    fbreak; };
    "}"           { EMITS(RCURLY)
                    top--
                    fnext *stack[top];
                    fbreak; };
    "{"           { EMITS(LCURLY)
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext inside_dollar_curly;
                    fbreak; };

    "''" (' '* ('\n' & NL))?
                  { EMITS(IND_STRING_OPEN);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext ind_string;
                    fbreak; };

    PATH          { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    HPATH         { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    SPATH         { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    URI           { EMIT_TEXT(URI, ts, te, data[ts:te])
                    fbreak; };

    [ \t\r\n]+ & NL;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT)
                    fbreak; };

    '"'           { EMITS(DQUOTE);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext string;
                    fbreak; };

    "="           { EMITS(EQS)
                    fbreak; };
    "?"           { EMITS(QMARK)
                    fbreak; };
    ","           { EMITS(COMMA)
                    fbreak; };
    "("           { EMITS(LPAREN)
                    fbreak; };
    ")"           { EMITS(RPAREN)
                    fbreak; };
    "["           { EMITS(LBRACK)
                    fbreak; };
    "]"           { EMITS(RBRACK)
                    fbreak; };
  *|;

  main := |*
    "if"          { EMITS(IF)
                    fbreak; };
    "then"        { EMITS(THEN)
                    fbreak; };
    "else"        { EMITS(ELSE)
                    fbreak; };
    "assert"      { EMITS(ASSERT)
                    fbreak; };
    "with"        { EMITS(WITH)
                    fbreak; };
    "let"         { EMITS(LET)
                    fbreak; };
    "in"          { EMITS(IN)
                    fbreak; };
    "rec"         { EMITS(REC)
                    fbreak; };
    "inherit"     { EMITS(INHERIT)
                    fbreak; };
    "or"          { EMITS(OR_KW)
                    fbreak; };
    "..."         { EMITS(ELLIPSIS)
                    fbreak; };
    ":"           { EMITS(COLON)
                    fbreak; };
    ";"           { EMITS(SCOLON)
                    fbreak; };
    "@"           { EMITS(AT)
                    fbreak; };
    "-"           { EMITS(MINUS)
                    fbreak; };
    "+"           { EMITS(PLUS)
                    fbreak; };
    "!"           { EMITS(NOT)
                    fbreak; };
    "*"           { EMITS(STAR)
                    fbreak; };

    "=="          { EMITS(EQ)
                    fbreak; };
    "!="          { EMITS(NEQ)
                    fbreak; };
    "<="          { EMITS(LEQ)
                    fbreak; };
    ">="          { EMITS(GEQ)
                    fbreak; };
    "&&"          { EMITS(AND)
                    fbreak; };
    "||"          { EMITS(OR)
                    fbreak; };
    "->"          { EMITS(IMPL)
                    fbreak; };
    "//"          { EMITS(UPDATE)
                    fbreak; };
    "++"          { EMITS(CONCAT)
                    fbreak; };

    ID            { EMIT_TEXT(ID, ts, te, data[ts:te])
                    fbreak; };
    INT           { EMIT_TEXT(INT, ts, te, data[ts:te])
                    fbreak; };
    FLOAT         { EMIT_TEXT(FLOAT, ts, te, data[ts:te])
                    fbreak; };

    "${"          { EMITS(DOLLAR_CURLY);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext inside_dollar_curly;
                    fbreak; };
    "}"           { EMITS(RCURLY)
                    fbreak; };
    "{"           { EMITS(LCURLY)
                    fbreak; };

    "''" (' '* ('\n' & NL))?
                  { EMITS(IND_STRING_OPEN);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext ind_string;
                    fbreak; };

    PATH          { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    HPATH         { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    SPATH         { EMIT_TEXT(PATH, ts, te, data[ts:te])
                    fbreak; };
    URI           { EMIT_TEXT(URI, ts, te, data[ts:te])
                    fbreak; };

    [ \t\r\n]+ & NL;

    SCOMMENT;
    LCOMMENT;

    "."           { EMITS(DOT)
                    fbreak; };

    '"'           { EMITS(DQUOTE);
                    if len(stack) <= top {
                      stack = append(stack, 0)
                    }
                    stack[top] = ftargs
                    top++
                    fnext string;
                    fbreak; };

    "="           { EMITS(EQS)
                    fbreak; };
    "?"           { EMITS(QMARK)
                    fbreak; };
    ","           { EMITS(COMMA)
                    fbreak; };
    "("           { EMITS(LPAREN)
                    fbreak; };
    ")"           { EMITS(RPAREN)
                    fbreak; };
    "["           { EMITS(LBRACK)
                    fbreak; };
    "]"           { EMITS(RBRACK)
                    fbreak; };
  *|;

  prepush {
    if len(stack) <= top {
      stack = append(stack, 0)
    }
  }
}%%