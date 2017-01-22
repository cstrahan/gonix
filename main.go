package main

import "log"

// var ex = []byte(". } http://foo.com <foo> 123 ~/a in 123.1 inter inherit ++ -> ${ ./foo/bar /** **/ ./baz # asdf\n foobar")
var ex = []byte(`"abc ${foo.bar {}}  baz" ''foo ''$ baz''`)

// var ex = []byte(`'' ''\n ''`)

// var ex = []byte(`'' ''$ ''`)

// var ex = []byte("foo/bar/baz")
// var ex = []byte("123.45 2.")

// var ex = []byte("./foo/bar ./quux")

//var ex = []byte(`foo: bar`)

func main() {
	log.Println("LEXING")
	tokens := lex(ex)
	log.Println(tokens)
}

type parser struct {
	pos    int
	tokens []Token
	data   []byte
}

func (self parser) tok(offset int) Token {
	return self.tokens[self.pos+offset]
}

func (self parser) advance(offset int) {
	self.pos += offset
}

type UnexpectedToken struct {
}

func (self UnexpectedToken) Error() string {
	return "Unexpected Token"
}

func (self parser) mustMatch(typ TokenType, offset int) (Token, error) {
	tok := self.tok(offset)
	if tok.TokenType != typ {
		return tok, UnexpectedToken{}
	}
	return tok, nil
}

func (self parser) parseExpr() (Expr, error) {
	return self.parseExprFunction()
}

func (self parser) parseExprFunction() (Expr, error) {
	t1 := self.tok(0)

	switch t1.TokenType {
	case ID:
		t2 := self.tok(1)

		switch t2.TokenType {
		case COLON:
			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			return ExprLambda{t1.Pos, t1.Text, false, []Formal{}, false, body}, nil
		case AT:
			_, err := self.mustMatch(LCURLY, 0)
			if err != nil {
				return nil, err
			}

			self.advance(3)

			formals, elipsis, err := self.parseFormals()
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(COLON, 1)
			if err != nil {
				return nil, err
			}

			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			return ExprLambda{t1.Pos, t1.Text, true, formals, elipsis, body}, nil
		}
	case ASSERT:
		self.advance(1)

		cond, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(SCOLON, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return ExprAssert{t1.Pos, cond, body}, nil
	case WITH:
		self.advance(1)

		attrs, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(SCOLON, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return ExprAssert{t1.Pos, attrs, body}, nil
	case LET:
		self.advance(1)

		attrs, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(IN, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return ExprLet{t1.Pos, attrs, body}, nil
	case LCURLY:
		self.advance(1)

		// TODO: fallback to attrs parsing here
		formals, elipsis, err := self.parseFormals()
		if err != nil {
			//return nil, err
			return self.parseExprIf()
		}

		_, err = self.mustMatch(RCURLY, 0)

		t2 := self.tok(1)
		switch t2.TokenType {
		case AT:
			nameTok, err := self.mustMatch(ID, 2)
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(COLON, 3)
			if err != nil {
				return nil, err
			}

			self.advance(3)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			return ExprLambda{t1.Pos, nameTok.Text, true, formals, elipsis, body}, nil
		case COLON:
			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			return ExprLambda{t1.Pos, []byte{}, true, formals, elipsis, body}, nil
		}
	default:
		return self.parseExprIf()
	}

	panic("unreachable")
}

func undefined() Expr {
	return nil
}

func (self parser) parseFormals() ([]Formal, bool, error) {
	return []Formal{}, false, nil
}

func (self parser) parseExprIf() (Expr, error) {
	t1 := self.tok(0)
	if t1.TokenType == IF {
		self.advance(1)

		cond, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(THEN, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		then, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(ELSE, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		els, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		return ExprIf{t1.Pos, cond, then, els}, nil
	} else {
		return self.parseExprOp()
	}
}

// interesting cases:
//   ---> ! { a = 1; } ? a || true
func (self parser) parseExprOp() (Expr, error) {
	ops := []Token{}
	opn := 0
	exprs := []Expr{}
	exprn := 0
	haveEq := false
	haveImpl := false

	for {
		// gobble-up all unary ops
		exprs = append(exprs, nil)
		link := &exprs[exprn]
		exprn++
	UNARY:
		tok := self.tok(0)
		switch tok.TokenType {
		case NOT:
			self.advance(1)

			expr := ExprOpNot{tok.Pos, nil}
			exprs = append(exprs, expr)
			link = &expr.Expr

			goto UNARY
		case MINUS:
			self.advance(1)

			expr := ExprApp{tok.Pos, []byte("__sub"), nil}
			exprs = append(exprs, expr)
			link = &expr.Arg

			goto UNARY
		}

		// parse expr
		expr, err := self.parseExprApp()
		if err != nil {
			return nil, err
		}
		*link = expr

		// next op
		// if this is the second occurence of a non-associative operator, bail.
		tok = self.tok(0)
		switch tok.TokenType {
		case EQ:
			if haveEq {
				return nil, UnexpectedToken{}
			}
			haveEq = true
		case NEQ:
			if haveEq {
				return nil, UnexpectedToken{}
			}
			haveEq = true
		case '<':
		case LEQ:
		case '>':
		case GEQ:
		case AND:
		case OR:
		case IMPL:
			if haveImpl {
				return nil, UnexpectedToken{}
			}
			haveImpl = true
		case UPDATE:
		case '?':
		case '+':
		case '-':
		case '*':
		case '/':
		case CONCAT:
		default:
			// we've consumed the last op; break
			goto SHUNT
		}

		// otherwise, push op and continue
		ops = append(ops, tok)
		opn++
	}

SHUNT:
	// shunting-yard

	// Associativity:
	// %nonassoc IMPL
	// %left OR
	// %left AND
	// %nonassoc EQ NEQ
	// %left '<' '>' LEQ GEQ
	// %right UPDATE
	// %left NOT
	// %left '+' '-'
	// %left '*' '/'
	// %right CONCAT
	// %nonassoc '?'
	// %nonassoc NEGATE
	for {
		if opn == 0 {
			return exprs[0], nil
		}

		opn--
		op1 := ops[opn]

		for opn > 0 {
			op2 := ops[opn-1]

			// all cases where:
			//   * op1 is left-associative and its precedence is less than or equal to that of op2, or
			//   * op1 is right associative, and has precedence less than that of op2
			switch op1.TokenType {
			case EQ:
			case NEQ:
			case '<':
			case LEQ:
			case '>':
			case GEQ:
			case AND:
			case OR:
			case IMPL:
			case UPDATE:
			case '?':
			case '+':
			case '-':
			case '*':
			case '/':
			case CONCAT:
			}

			switch op2.TokenType {
			case EQ:
			case NEQ:
			case '<':
			case LEQ:
			case '>':
			case GEQ:
			case AND:
			case OR:
			case IMPL:
			case UPDATE:
			case '?':
			case '+':
			case '-':
			case '*':
			case '/':
			case CONCAT:
			default:
				// we've consumed the last op; break
				goto SHUNT
			}

			//opn--
		}

	}

	return undefined(), nil
}

func (self parser) parseExprApp() (Expr, error) {
	return undefined(), nil
}

type Symbol []byte

type Expr interface{}

type Formal struct {
	Name Symbol
}

type ExprLambda struct {
	Pos
	Name       Symbol
	MatchAttrs bool
	Formals    []Formal
	Elipsis    bool
	Body       Expr
}

type ExprAssert struct {
	Pos
	Cond Expr
	Body Expr
}

type ExprWith struct {
	Pos
	Attrs Expr
	Body  Expr
}

type ExprLet struct {
	Pos
	Attrs Expr
	Body  Expr
}

type ExprIf struct {
	Pos
	Cond Expr
	Then Expr
	Else Expr
}

type ExprOpNot struct {
	Pos
	Expr Expr
}

type ExprApp struct {
	Pos
	Fun Expr
	Arg Expr
}

type ExpVar struct {
	Pos
	Name Symbol
}

//--------------------------------

//type Position struct {
//	Filename string // filename, if any
//	Offset   int    // offset, starting at 0
//	Line     int    // line number, starting at 1
//	Column   int    // column number, starting at 1 (byte count)
//}
// IsValid reports whether the position is valid.
//func (pos *Position) IsValid() bool { return pos.Line > 0 }

//type Token struct {
//	TokenType TokenType
//	Start     int
//	End       int
//	Text      []byte
//}

//--------------------------------
// EXPRS
