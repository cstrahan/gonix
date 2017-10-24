package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// var ex = []byte(". } http://foo.com <foo> 123 ~/a in 123.1 inter inherit ++ -> ${ ./foo/bar /** **/ ./baz # asdf\n foobar")

//var ex = []byte("1 + 2 * 3")

//var ex = []byte("{ a, b, a, c }@y: x")
//var ex = []byte("let x = 123; y = 321; in x")
//var ex = []byte("x.a or x.b or x.c")

//var ex = []byte("rec { a = 123; }")

//var ex = []byte("foo.\"abc${bar}def\" ")

//var ex = []byte(`"abc ${foo.bar {}}  baz" ''foo ''$ baz''`)
var ex = []byte(`''
  foo
''`)

//var ex = []byte(`'' ''\n ''`)

// var ex = []byte(`'' ''$ ''`)

// var ex = []byte("foo/bar/baz")
// var ex = []byte("123.45 2.")

// var ex = []byte("./foo/bar ./quux")

//var ex = []byte(`foo: bar`)

func printTok(tok Token) string {
	return fmt.Sprintf("%v:%v  %s  [%s]", tok.Pos.Start, tok.Pos.End, tok.TokenType, string(tok.Text))
}

//func (self Pos) String() string {
//	return fmt.Sprintf("(%v:%v)", self.Start, self.End)
//}

func main() {
	data := ex

	log.Println("LEXING")
	tokens := lex(ex)
	for _, tok := range tokens {
		fmt.Println(printTok(tok))
	}

	// synthesize some EOF tokens for the sake of lookahead
	eof := Token{TokenType: EOF}
	tokens = append(tokens, eof, eof, eof, eof)

	p := parser{pos: 0, tokens: tokens, data: data}
	expr, err := p.parseExpr()
	if err != nil {
		log.Fatalln("parse failed:", err)
	}

	spew.Dump(expr)
	//log.Println(expr)
}

type parser struct {
	pos    int
	tokens []Token
	data   []byte
}

func showAttrPath(attrPath []AttrName) string {
	out := ""
	first := true
	for _, i := range attrPath {
		if !first {
			out = out + "."
		} else {
			first = false
		}

		if i.Expr == nil {
			out = out + string(i.Symbol)
		} else {
			out = out + "\"${" + "<EXPR>" + "}" // TODO
			//out = out + "\"${" + i.Expr + "}" // TODO
		}
	}

	return out
}

func dupAttrPath(attrPath []AttrName, pos Pos, prevPos Pos) error {
	return fmt.Errorf("attribute ‘%s’ at %s already defined at %s", showAttrPath(attrPath), pos, prevPos)
}

func dupAttrSym(attr Symbol, pos Pos, prevPos Pos) error {
	return fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr, pos, prevPos)
}

func addAttr(attrs ExprAttrs, attrPath []AttrName, expr Expr, pos Pos) error {
	for _, attr := range attrPath[:len(attrPath)-1] {
		if attr.Expr == nil {
			if existing, ok := attrs.AttrDefs[string(attr.Symbol)]; ok {
				if !existing.Inherited {
					if prevAttrs, ok := existing.Expr.(ExprAttrs); ok {
						attrs = prevAttrs
					} else {
						return dupAttrPath(attrPath, pos, existing.Pos)
					}
				} else {
					return dupAttrPath(attrPath, pos, existing.Pos)
				}
			} else {
				nested := ExprAttrs{
					AttrDefs:        map[string]AttrDef{},
					DynamicAttrDefs: map[string]DynamicAttrDef{},
				}
				attrs.AttrDefs[string(attr.Symbol)] = AttrDef{
					Expr:      nested,
					Inherited: false,
				}
				attrs = nested
			}
		} else {
			nested := ExprAttrs{
				AttrDefs:        map[string]AttrDef{},
				DynamicAttrDefs: map[string]DynamicAttrDef{},
			}

			attrs.DynamicAttrDefs[string(attr.Symbol)] = DynamicAttrDef{
				NameExpr:  attr.Expr,
				ValueExpr: nested,
			}
			attrs = nested
		}
	}

	attr := attrPath[len(attrPath)-1]
	if attr.Expr == nil {
		if existing, ok := attrs.AttrDefs[string(attr.Symbol)]; ok {
			_ = existing
			return dupAttrPath(attrPath, pos, existing.Pos)
		} else {
			attrs.AttrDefs[string(attr.Symbol)] = AttrDef{
				Expr:      expr,
				Inherited: false,
			}
		}
	}

	return nil
}

func addFormal(pos Pos, formals *Formals, formal Formal) error {
	if _, ok := formals.ArgNames[string(formal.Name)]; ok {
		return fmt.Errorf("duplicate formal function argument ‘%s’ at %s", formal.Name, pos)
	}

	formals.ArgNames[string(formal.Name)] = struct{}{}
	formals.Formals = append(formals.Formals, formal)
	return nil
}

func stripIndentation(pos Pos, exprs []Expr) Expr {
	if len(exprs) == 0 {
		return ExprString{[]byte("")}
	}

	atStartOfLine := true
	minIndent := 1000000
	curIndent := 0

	// Figure out minimum indentation.
	for _, expr := range exprs {
		e, ok := expr.(exprIndStr)
		if !ok {
			// Anti-quotations end the current start-of-line whitespace.
			if atStartOfLine {
				atStartOfLine = false
				if curIndent < minIndent {
					minIndent = curIndent
				}
			}
			continue
		}
		for j := 0; j < len(e.String); j++ {
			if atStartOfLine {
				if e.String[j] == ' ' {
					curIndent++
				} else if e.String[j] == '\n' {
					curIndent = 0
				} else {
					atStartOfLine = false
					if curIndent < minIndent {
						minIndent = curIndent
					}
				}
			} else if e.String[j] == '\n' {
				atStartOfLine = true
				curIndent = 0
			}
		}
	}

	// Strip spaces from each line.
	exprs2 := []Expr{}
	atStartOfLine = true
	curDropped := 0
	for n, i := range exprs {
		e, ok := i.(exprIndStr)
		if !ok {
			atStartOfLine = false
			curDropped = 0
			exprs2 = append(exprs2, i)
			continue
		}

		str2 := ""
		for j := 0; j < len(e.String); j++ {
			if atStartOfLine {
				if e.String[j] == ' ' {
					if curDropped >= minIndent {
						str2 += string(e.String[j])
					}
					curDropped++
				} else if e.String[j] == '\n' {
					curDropped = 0
					str2 += string(e.String[j])
				} else {
					atStartOfLine = false
					curDropped = 0
					str2 += string(e.String[j])
				}
			} else {
				str2 += string(e.String[j])
				if e.String[j] == '\n' {
					atStartOfLine = true
				}
			}
		}

		// Remove the last line if it is empty and consists only of spaces.
		if n == len(exprs)-1 {
			idx := strings.LastIndexByte(str2, '\n')
			if idx != -1 {
				for k := idx + 1; k < len(str2); k++ {
					if str2[k] != ' ' {
						goto Done
					}
				}

				str2 = str2[0:idx]
			}
		}

	Done:
		exprs2 = append(exprs2, ExprString{String: []byte(str2)})
	}

	// If this is a single string, then don't do a concatenation.
	if len(exprs2) == 1 {
		if e, ok := exprs2[0].(ExprString); ok {
			return e
		}
	}

	return ExprConcatStrings{Exprs: exprs2, ForceString: true}
}

func (self *parser) tok(offset int) Token {
	return self.tokens[self.pos+offset]
}

func (self *parser) advance(offset int) {
	self.pos += offset
}

func unexpected() error {
	panic("unexpected token")
	return UnexpectedToken{}
}

type UnexpectedToken struct {
}

func (self UnexpectedToken) Error() string {
	return "Unexpected Token"
}

func (self *parser) mustMatch(typ TokenType, offset int) (Token, error) {
	tok := self.tok(offset)
	if tok.TokenType != typ {
		return tok, unexpected()
	}
	return tok, nil
}

func (self *parser) parseExpr() (Expr, error) {
	return self.parseExprFunction()
}

func (self *parser) parseExprFunction() (Expr, error) {
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

			lambda := ExprLambda{
				Pos:        t1.Pos,
				Name:       t1.Text,
				MatchAttrs: false,
				Formals:    Formals{Formals: []Formal{}, ArgNames: map[string]struct{}{}},
				Body:       body,
			}

			return lambda, nil
		case AT:
			_, err := self.mustMatch(LCURLY, 2)
			if err != nil {
				return nil, err
			}

			self.advance(3)

			formals, err := self.parseFormals()
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

			lambda := ExprLambda{
				Pos:        t1.Pos,
				Name:       t1.Text,
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
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

		return ExprWith{t1.Pos, attrs, body}, nil
	case LET:
		self.advance(1)

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		if len(binds.DynamicAttrDefs) != 0 {
			return nil, fmt.Errorf("dynamic attributes not allowed in let at %s", t1)
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

		return ExprLet{t1.Pos, binds, body}, nil
	case LCURLY:
		if !self.isFunctionWithFormals() {
			goto Fallthrough
		}

		self.advance(1)

		formals, err := self.parseFormals()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)

		t2 := self.tok(1)
		switch t2.TokenType {
		case AT:
			nameTok, err := self.mustMatch(ID, 2)
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(':', 3)
			if err != nil {
				return nil, err
			}

			self.advance(4)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := ExprLambda{
				Pos:        t1.Pos,
				Name:       nameTok.Text,
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
		case COLON:
			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := ExprLambda{
				Pos:        t1.Pos,
				Name:       []byte{},
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
		}
	default:
		goto Fallthrough
	}

Fallthrough:
	return self.parseExprIf()
}

func (self *parser) isTok(offset int, typ TokenType) bool {
	return self.tok(offset).TokenType == typ
}

// Use this to disambiguate a function with formal arguments from an attr-set.
//
// The subsequent tokens are part of a function with formals when they are one
// of the following sequences:
//
//     '{' ID '}'
//     '{' ID ','
//     '{' ID '?'
//     '{' '}' ':'
//     '{' '}' '@'
//     '{' ELLIPSIS
//
// Otherwise, the tokens begin an attr-set.
func (self *parser) isFunctionWithFormals() bool {
	if !self.isTok(0, '{') {
		return false
	}

	return (self.isTok(1, ID) && self.isTok(2, '}')) ||
		(self.isTok(1, ID) && self.isTok(2, ',')) ||
		(self.isTok(1, ID) && self.isTok(2, '?')) ||
		(self.isTok(1, '}') && self.isTok(2, ':')) ||
		(self.isTok(1, '}') && self.isTok(2, '@')) ||
		self.isTok(1, ELLIPSIS)
}

func undefined() Expr {
	return nil
}

func (self *parser) parseExprIf() (Expr, error) {
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
func (self *parser) parseExprOp() (Expr, error) {
	ops := []Token{}
	exprs := []Expr{}
	haveEq := false
	haveImpl := false

	for {
		// push all unary operators
		for {
			tok := self.tok(0)
			if tok.TokenType == NOT {
				ops = append(ops, tok)
				self.advance(1)
			} else if tok.TokenType == MINUS {
				tok.TokenType = NEGATE
				ops = append(ops, tok)
				self.advance(1)
			} else {
				break
			}
		}

		// parse/push expr
		expr, err := self.parseExprApp()
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)

		// next op
		op1 := self.tok(0)
		if isOperator(op1.TokenType) {
			self.advance(1)
			op1_ := operators[op1.TokenType]

			// sanity check
			switch op1.TokenType {
			case EQ:
				if haveEq {
					return nil, unexpected()
				}
				haveEq = true
			case NEQ:
				if haveEq {
					return nil, unexpected()
				}
				haveEq = true
			case IMPL:
				if haveImpl {
					return nil, unexpected()
				}
				haveImpl = true
			}

			for len(ops) > 0 {
				op2 := ops[len(ops)-1]
				op2_ := operators[op2.TokenType]
				if op1_.Assoc == ASSOC_LEFT && op1_.Prec <= op2_.Prec ||
					op1_.Assoc == ASSOC_RIGHT && op1_.Prec < op2_.Prec {
					ops = ops[:len(ops)-1] // pop op2 off stack
					// TODO: produce/push expr
					switch op2.TokenType {
					case '!':
						exprs[len(exprs)-1] = ExprOpNot{op2.Pos, exprs[len(exprs)-1]}
					case NEGATE:
						exprs[len(exprs)-1] = ExprOpNegate{op2.Pos, exprs[len(exprs)-1]}
					case EQ:
						exprs[len(exprs)-2] = ExprOpEQ{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case NEQ:
						exprs[len(exprs)-2] = ExprOpNEQ{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '<':
						exprs[len(exprs)-2] = ExprOpLE{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case LEQ:
						exprs[len(exprs)-2] = ExprOpLEQ{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '>':
						exprs[len(exprs)-2] = ExprOpGE{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case GEQ:
						exprs[len(exprs)-2] = ExprOpGEQ{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case AND:
						exprs[len(exprs)-2] = ExprOpAnd{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case OR:
						exprs[len(exprs)-2] = ExprOpOr{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case IMPL:
						exprs[len(exprs)-2] = ExprOpImpl{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case UPDATE:
						exprs[len(exprs)-2] = ExprOpUpdate{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '?':
						exprs[len(exprs)-2] = ExprOpHasAttr{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '+':
						exprs[len(exprs)-2] = ExprOpAdd{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '-':
						exprs[len(exprs)-2] = ExprOpSub{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '*':
						exprs[len(exprs)-2] = ExprOpMult{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case '/':
						exprs[len(exprs)-2] = ExprOpDiv{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					case CONCAT:
						exprs[len(exprs)-2] = ExprOpConcat{op2.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					}
				} else {
					break
				}
			}

			// push the new op to the stack
			ops = append(ops, op1)
		} else {
			// no more operators/exprs; pop remaining ops from stack
			// TODO: find good way to dedupe this
			for len(ops) > 0 {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1] // pop op off stack
				switch op.TokenType {
				case '!':
					exprs[len(exprs)-1] = ExprOpNot{op.Pos, exprs[len(exprs)-1]}
				case NEGATE:
					exprs[len(exprs)-1] = ExprOpNegate{op.Pos, exprs[len(exprs)-1]}
				case EQ:
					exprs[len(exprs)-2] = ExprOpEQ{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case NEQ:
					exprs[len(exprs)-2] = ExprOpNEQ{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '<':
					exprs[len(exprs)-2] = ExprOpLE{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case LEQ:
					exprs[len(exprs)-2] = ExprOpLEQ{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '>':
					exprs[len(exprs)-2] = ExprOpGE{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case GEQ:
					exprs[len(exprs)-2] = ExprOpGEQ{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case AND:
					exprs[len(exprs)-2] = ExprOpAnd{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case OR:
					exprs[len(exprs)-2] = ExprOpOr{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case IMPL:
					exprs[len(exprs)-2] = ExprOpImpl{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case UPDATE:
					exprs[len(exprs)-2] = ExprOpUpdate{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '?':
					exprs[len(exprs)-2] = ExprOpHasAttr{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '+':
					exprs[len(exprs)-2] = ExprOpAdd{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '-':
					exprs[len(exprs)-2] = ExprOpSub{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '*':
					exprs[len(exprs)-2] = ExprOpMult{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case '/':
					exprs[len(exprs)-2] = ExprOpDiv{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				case CONCAT:
					exprs[len(exprs)-2] = ExprOpConcat{op.Pos, exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				}
			}

			return exprs[0], nil
		}

	}

	return exprs[0], nil
}

const (
	ASSOC_LEFT  = 0
	ASSOC_RIGHT = 1
	ASSOC_NONE  = 2
)

type Op struct {
	Assoc int
	Prec  int
}

var operators = [...]Op{
	IMPL:   Op{Assoc: ASSOC_NONE, Prec: 0},
	OR:     Op{Assoc: ASSOC_LEFT, Prec: 1},
	AND:    Op{Assoc: ASSOC_LEFT, Prec: 2},
	EQ:     Op{Assoc: ASSOC_NONE, Prec: 3},
	NEQ:    Op{Assoc: ASSOC_NONE, Prec: 3},
	'<':    Op{Assoc: ASSOC_LEFT, Prec: 4},
	'>':    Op{Assoc: ASSOC_LEFT, Prec: 4},
	LEQ:    Op{Assoc: ASSOC_LEFT, Prec: 4},
	GEQ:    Op{Assoc: ASSOC_LEFT, Prec: 4},
	UPDATE: Op{Assoc: ASSOC_RIGHT, Prec: 5},
	NOT:    Op{Assoc: ASSOC_LEFT, Prec: 6},
	'+':    Op{Assoc: ASSOC_LEFT, Prec: 7},
	'-':    Op{Assoc: ASSOC_LEFT, Prec: 7},
	'*':    Op{Assoc: ASSOC_LEFT, Prec: 8},
	'/':    Op{Assoc: ASSOC_LEFT, Prec: 8},
	CONCAT: Op{Assoc: ASSOC_RIGHT, Prec: 9},
	'?':    Op{Assoc: ASSOC_NONE, Prec: 10},
	NEGATE: Op{Assoc: ASSOC_NONE, Prec: 11},
}

func isOperator(ty TokenType) bool {
	return ty == IMPL ||
		ty == OR ||
		ty == AND ||
		ty == EQ ||
		ty == NEQ ||
		ty == '<' ||
		ty == '>' ||
		ty == LEQ ||
		ty == GEQ ||
		ty == UPDATE ||
		ty == NOT ||
		ty == '+' ||
		ty == '-' ||
		ty == '*' ||
		ty == '/' ||
		ty == CONCAT ||
		ty == '?' ||
		ty == NEGATE
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (self *parser) parseExprApp() (Expr, error) {
	expr, err := self.parseExprSelect()
	if err != nil {
		return nil, err
	}

	for {
		switch self.tok(0).TokenType {
		case ID, INT, FLOAT, '"', IND_STRING_OPEN, PATH, HPATH, SPATH, URI, '(', LET, REC, '{', '[':

			body, err := self.parseExprSelect()
			if err != nil {
				return nil, err
			}

			expr = ExprApp{Pos{}, expr, body}
		default:
			return expr, nil
		}
	}

	return expr, nil
}

func (self *parser) parseExprSelect() (Expr, error) {
	simple, err := self.parseExprSimple()
	if err != nil {
		return nil, err
	}

	tok := self.tok(0)
	switch tok.TokenType {
	case DOT:
		self.advance(1)
		path, err := self.parseAttrPath()
		if err != nil {
			return nil, err
		}

		var def Expr
		tok := self.tok(0)
		if tok.TokenType == OR_KW {
			self.advance(1)
			def, err = self.parseExprSelect()
			if err != nil {
				return nil, err
			}
		}

		return ExprSelect{Pos{}, simple, path, def}, nil
	case OR_KW:
		self.advance(1)
		return ExprApp{Pos{}, simple, ExprVar{tok.Pos, []byte("or")}}, nil
	}

	return simple, nil
}

func (self *parser) parseExprSimple() (Expr, error) {
	t1 := self.tok(0)
	switch t1.TokenType {
	case ID:
		self.advance(1)
		return ExprVar{t1.Pos, t1.Text}, nil
	case INT:
		self.advance(1)

		n, err := strconv.ParseInt(string(t1.Text), 10, 32)
		if err != nil {
			return nil, err
		}

		return ExprInt{int(n)}, nil
	// case FLOAT:  // TODO
	case '"':
		self.advance(1)

		stringParts, err := self.parseStringParts()
		if err != nil {
			return nil, err
		}
		_, err = self.mustMatch('"', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)
		return stringParts, nil
	case IND_STRING_OPEN:
		self.advance(1)

		stringParts, err := self.parseIndStringParts()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(IND_STRING_CLOSE, 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return stripIndentation(t1.Pos, stringParts), nil
	case PATH:
		self.advance(1)
		return ExprPath{string(t1.Text)}, nil
	case HPATH:
		self.advance(1)
		return ExprPath{string(t1.Text)}, nil
	case SPATH:
		self.advance(1)
		return ExprPath{string(t1.Text)}, nil
	case URI:
		self.advance(1)
		return ExprString{t1.Text}, nil
	case '(':
		self.advance(1)

		expr, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(')', 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)
		return expr, nil
	case LET:
		self.advance(1)
		_, err := self.mustMatch('{', 0)
		if err != nil {
			return nil, err
		}

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds.Recursive = true
		attrPath := []AttrName{
			AttrName{Symbol: []byte("body")},
		}
		return ExprSelect{Pos{}, binds, attrPath, nil}, nil
	case REC:
		self.advance(1)

		_, err := self.mustMatch('{', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds.Recursive = true
		return binds, nil
	case '{':
		self.advance(1)
		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return binds, err
	case '[':
		self.advance(1)
		list, err := self.parseExprList()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(']', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return list, err
	}

	return nil, unexpected()
}

func (self *parser) parseStringParts() (Expr, error) {
	parts := []Expr{}

	tok := self.tok(0)

	// Exit early if this is just a simple string literal.
	// We'll make use of the fact that we have an ExprString in this case
	// elsewhere.
	if t2 := self.tok(1); tok.TokenType == STR && t2.TokenType == '"' {
		self.advance(1)
		str := ExprString{[]byte(tok.Text)}
		return str, nil
	}

	for {
		switch tok.TokenType {
		case STR:
			self.advance(1)

			str := ExprString{tok.Text}
			parts = append(parts, str)
		case DOLLAR_CURLY:
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return nil, err
			}
			parts = append(parts, expr)

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}
			self.advance(1)
		default:
			goto Break
		}

		tok = self.tok(0)
	}

Break:
	concatExpr := ExprConcatStrings{Exprs: parts, ForceString: true}
	return concatExpr, nil
}

func (self *parser) parseIndStringParts() ([]Expr, error) {
	parts := []Expr{}

	tok := self.tok(0)

	for {
		switch tok.TokenType {
		case IND_STR:
			self.advance(1)

			str := exprIndStr{String: tok.Text}
			parts = append(parts, str)
		case DOLLAR_CURLY:
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return nil, err
			}
			parts = append(parts, expr)

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}
			self.advance(1)
		default:
			goto Break
		}

		tok = self.tok(0)
	}

Break:
	return parts, nil
}

func (self *parser) parseBinds() (ExprAttrs, error) {
	binds := ExprAttrs{
		AttrDefs:        map[string]AttrDef{},
		DynamicAttrDefs: map[string]DynamicAttrDef{},
	}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case '}', IN:
			return binds, nil
		}

		switch tok.TokenType {
		case INHERIT:
			self.advance(1)
			tok = self.tok(0)

			switch tok.TokenType {
			case '(':
				self.advance(1)
				tok = self.tok(0)

				expr, err := self.parseExpr()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(')', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				attrs, err := self.parseAttrs()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(';', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				for _, attr := range attrs {
					if _, ok := binds.AttrDefs[string(attr.Symbol)]; ok {
						return binds, fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr.Symbol, "???", "???")
					}

					binds.AttrDefs[string(attr.Symbol)] = AttrDef{
						Expr: ExprSelect{
							Expr:     expr,
							AttrPath: []AttrName{AttrName{Symbol: attr.Symbol}},
						},
						Inherited: false,
					}
				}
			default:
				attrs, err := self.parseAttrs()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(';', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				for _, attr := range attrs {
					if _, ok := binds.AttrDefs[string(attr.Symbol)]; ok {
						return binds, fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr.Symbol, "???", "???")
					}

					binds.AttrDefs[string(attr.Symbol)] = AttrDef{
						Expr: ExprVar{
							Name: attr.Symbol,
						},
						Inherited: true,
					}
				}
			}
		default:
			attrpath, err := self.parseAttrPath()
			if err != nil {
				return binds, err
			}

			_, err = self.mustMatch('=', 0)
			if err != nil {
				return binds, err
			}
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return binds, err
			}

			_, err = self.mustMatch(';', 0)
			if err != nil {
				return binds, err
			}
			self.advance(1)

			err = addAttr(binds, attrpath, expr, Pos{})
			if err != nil {
				return binds, err
			}
		}
	}

	return binds, nil
}

func (self *parser) parseAttrs() ([]AttrName, error) {
	attrs := []AttrName{}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case ID, OR_KW:
			attr, err := self.parseAttr()
			if err != nil {
				return nil, err
			}

			attrs = append(attrs, AttrName{Symbol: Symbol(attr)})
		case '"', DOLLAR_CURLY:
			strExpr, err := self.parseStringAttr()
			if err != nil {
				return nil, err
			}

			if str, ok := strExpr.(ExprString); ok {
				attrs = append(attrs, AttrName{Symbol: Symbol(str.String)})
			} else {
				return nil, fmt.Errorf("dynamic attributes not allowed in inherit at %s", "???")
			}
		default:
			return attrs, nil
		}
	}

	return attrs, nil
}

func (self *parser) parseAttrPath() ([]AttrName, error) {
	attrs := []AttrName{}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case ID, OR_KW:
			attr, err := self.parseAttr()
			if err != nil {
				return nil, err
			}

			attrs = append(attrs, AttrName{Symbol: Symbol(attr)})
		case '"', DOLLAR_CURLY:
			strExpr, err := self.parseStringAttr()
			if err != nil {
				return nil, err
			}

			if str, ok := strExpr.(ExprString); ok {
				attrs = append(attrs, AttrName{Symbol: Symbol(str.String)})
			} else {
				attrs = append(attrs, AttrName{Expr: strExpr})
			}
		}

		tok = self.tok(0)
		if tok.TokenType != '.' {
			break
		}
		self.advance(1)
	}

	return attrs, nil
}

func (self *parser) parseAttr() (string, error) {
	tok := self.tok(0)
	self.advance(1)
	switch tok.TokenType {
	case ID:
		return string(tok.Text), nil
	case OR_KW:
		return "or", nil
	default:
		return "", unexpected()
	}
}

func (self *parser) parseStringAttr() (Expr, error) {
	tok := self.tok(0)
	self.advance(1)
	switch tok.TokenType {
	case '"':
		expr, err := self.parseStringParts()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('"', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return expr, err
	case DOLLAR_CURLY:
		expr, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return expr, err
	default:
		return nil, unexpected()
	}
}

func (self *parser) parseExprList() (ExprList, error) {
	elems := []Expr{}
	tok := self.tok(0)

	for {
		if tok.TokenType == ']' {
			break
		}

		elem, err := self.parseExprSelect()
		if err != nil {
			return ExprList{}, err
		}

		elems = append(elems, elem)

		tok = self.tok(0)
	}

	return ExprList{Elems: elems}, nil
}

func (self *parser) parseFormals() (Formals, error) {
	formals := Formals{
		ArgNames: map[string]struct{}{},
		Formals:  []Formal{},
		Ellipsis: false,
	}

	tok := self.tok(0)

	for {
		switch tok.TokenType {
		case ELLIPSIS:
			self.advance(1)
			formals.Ellipsis = true
			return formals, nil
		case ',':
			self.advance(1)
		case ID:
			formal, err := self.parseFormal()
			if err != nil {
				return Formals{}, err
			}

			err = addFormal(tok.Pos, &formals, formal)
			if err != nil {
				return Formals{}, err
			}
		default:
			return formals, nil
		}

		tok = self.tok(0)
	}

	return formals, nil
}

func (self *parser) parseFormal() (Formal, error) {
	tokId, err := self.mustMatch(ID, 0)
	if err != nil {
		return Formal{}, err
	}
	self.advance(1)

	if self.tok(0).TokenType == '?' {
		self.advance(1)
		expr, err := self.parseExpr()
		if err != nil {
			return Formal{}, err
		}

		return Formal{Name: tokId.Text, Def: expr}, nil
	} else {
		return Formal{Name: tokId.Text, Def: nil}, nil
	}
}

type Symbol []byte

type Expr interface{}

type Formals struct {
	Formals  []Formal
	ArgNames map[string]struct{}
	Ellipsis bool
}

type Formal struct {
	Name Symbol
	Def  Expr
}

type ExprConcatStrings struct {
	ForceString bool
	Exprs       []Expr
}

type ExprAttrs struct {
	Recursive       bool
	AttrDefs        map[string]AttrDef
	DynamicAttrDefs map[string]DynamicAttrDef
}

type AttrDef struct {
	Pos       Pos
	Inherited bool
	Expr      Expr
	//Displacement int
}

type DynamicAttrDef struct {
	//Pos Pos
	NameExpr  Expr
	ValueExpr Expr
}

type ExprLambda struct {
	Pos
	Name       Symbol
	MatchAttrs bool
	Formals    Formals
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
	Attrs ExprAttrs
	Body  Expr
}

type ExprIf struct {
	Pos
	Cond Expr
	Then Expr
	Else Expr
}

type ExprOpNegate struct {
	Pos
	Expr Expr
}

type ExprOpNEQ struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpEQ struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpLE struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpLEQ struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpGE struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpGEQ struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpAnd struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpOr struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpImpl struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpUpdate struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpHasAttr struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpAdd struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpSub struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpMult struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpDiv struct {
	Pos
	Expr1 Expr
	Expr2 Expr
}

type ExprOpConcat struct {
	Pos
	Expr1 Expr
	Expr2 Expr
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

type ExprVar struct {
	Pos
	Name Symbol
}

type ExprSelect struct {
	Pos
	Expr     Expr
	AttrPath []AttrName
	Def      Expr
}

type AttrName struct {
	Symbol Symbol
	Expr   Expr
}

type ExprInt struct {
	Val int
}

type ExprPath struct {
	Path string
}

type ExprString struct {
	String []byte
}

// Temporary structure during parse.
type exprIndStr struct {
	String []byte
}

type ExprList struct {
	Elems []Expr
}

type ExprPos struct {
	Pos Pos
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
