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
  if l.p == l.pe {
    return Token{ TokenType: TokenType(EOF), Pos: Pos { l.pe, l.pe, l.lineCount+1, (l.pe-l.lineStart)+1 }, Text: nil }, nil
  }

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

  token := Token{TokenType: EOF}

  var err error
%% write exec;

  // store state in l
  l.cs = cs
  l.p = p
  l.top = top
  l.lineStart = lineStart
  l.lineCount = lineCount
  l.act = act

  if cs == %%{ write error; }%% {
    err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
  } else if token.TokenType == EOF {
    return Token{ TokenType: TokenType(EOF), Pos: Pos { l.pe, l.pe, l.lineCount+1, (l.pe-l.lineStart)+1 }, Text: nil }, nil
  }

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
                  {
                    // don't consume the final " char.
                    // we'll set the next state back to string,
                    // which will handle emitting the final " char
                    // and popping the stack.
                    fhold;
                    EMIT_TEXT(STR, ts, te-1, data[ts:te-1]);
                    fnext string;
                    fbreak;
                  };
    ( [^$"\\] | '$' [^{"\\] | '\\' (any & NL) | '$\\' (any & NL) )+
                  { EMIT_TEXT(STR, ts, te, data[ts:te]); // ?
                    fbreak; };
    "${"          { EMITS(DOLLAR_CURLY);
                    if len(stack) <= top {
                      stack = append(stack, lexContext{})
                    }
                    stack[top] = lexContext{1, ftargs}
                    top++
                    fnext main;
                    fbreak;
                    };
    '"'           { EMITS(DQUOTE);
                    fnext main;
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
                            stack = append(stack, lexContext{})
                          }
                          stack[top] = lexContext{1, ftargs}
                          top++
                          fnext main;
                          fbreak; };
    "''"                { EMITS(IND_STRING_CLOSE);
                          fnext main;
                          fbreak; };
    "'"                 { EMIT_TEXT(IND_STR, ts, te, []byte("'"))
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
                    stack[top-1].curlyCount += 1
                    fbreak; };
    "}"           { EMITS(RCURLY)
                    stack[top-1].curlyCount -= 1
                    if stack[top-1].curlyCount == 0 {
                      fnext *stack[top-1].resumeState;
                      top--
                    }
                    fbreak; };
    "{"           { EMITS(LCURLY)
                    stack[top-1].curlyCount += 1
                    fbreak; };

    "''" (' '* ('\n' & NL))?
                  { EMITS(IND_STRING_OPEN);
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